package topcoder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	// Current version of go-topcoder.
	libraryVersion = "0.0.1"

	userAgent   = "go-topcoder/" + libraryVersion
	topcoderApi = "http://api.topcoder.com/v2/"

	contentType = "application/json"
)

type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent used when communicating with the Topcoder API.
	UserAgent string

	// Services used for talking to different parts of the Topcoder API.
	UserProfile *UserProfileApi
	Members     *MembersApi
	Data        *DataApi
}

// Kind of a singleton allowing users to make requests for public data without
// explicitly keeping a reference to a go-topcoder client
var (
	client      *Client
	UserProfile *UserProfileApi
	Members     *MembersApi
	Data        *DataApi
)

func init() {
	client = NewClient(nil)
	UserProfile = client.UserProfile
	Members = client.Members
	Data = client.Data
}

// urlParameters adds the parameters in params as URL query parameters to sUrl.
// params must be a struct whose fields may contain "url" tags.
func urlParameters(sUrl string, params interface{}) (string, error) {
	if p := reflect.ValueOf(params); p.Kind() == reflect.Ptr && p.IsNil() {
		return sUrl, nil
	}

	url, err := url.Parse(sUrl)
	if err != nil {
		return sUrl, err
	}

	queryString, err := query.Values(params)
	if err != nil {
		return sUrl, err
	}

	url.RawQuery = queryString.Encode()
	return url.String(), nil
}

// NewClient returns a new Topcoder API client. A http.DefaultClient is used
// when a nil httpClient is provided.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(topcoderApi)
	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.UserProfile = &UserProfileApi{client: c}
	c.Members = &MembersApi{client: c}
	c.Data = &DataApi{client: c}
	return c
}

// NewRequest creates an API request.
// A specific endpoint can be provided, and is resolved relative to the BaseURL
// of the Client. Endpoints should be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included as
// the request body.
func (c *Client) NewRequest(method, relPath string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}

	url := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", contentType)
	if len(c.UserAgent) > 0 {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	rawResponse, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	response := newResponse(rawResponse)

	err = CheckResponse(rawResponse)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, rawResponse.Body)
		} else {
			err = json.NewDecoder(rawResponse.Body).Decode(v)
			// ignore EOF errors caused by empty response body
			if err == io.EOF {
				err = nil
			}
		}
	}

	return response, err
}

// A Topcoder API response. It wraps the standard http.Response and might be
// useful later for Topcoder' specific response handling
type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// CheckResponse checks the API response for errors.
// A response is considered an error if its status code is outside the 200 range.
// API error responses might have a JSON response body that maps to ErrorResponse.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}

// Reports the errors caused by the Topcoder API request.
type ErrorResponse struct {
	Response *http.Response
	Err      struct {
		Message string `json:"description"`
		Details string `json:"details"`
	} `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Err.Message, r.Err.Details)
}
