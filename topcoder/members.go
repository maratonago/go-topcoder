package topcoder

import "fmt"

type MembersApi struct {
	client *Client
}

type TrackType string

const (
	Design      TrackType = "design"
	Development TrackType = "develop"
	DataScience TrackType = "data"
)

type TopTrackListOptions struct {
	ListOptions
}

type RankItem struct {
	Rank              uint   `json:"rank,omitempty"`
	Handle            string `json:"handle,omitempty"`
	UserId            uint64 `json:"userId,omitempty"`
	Color             string `json:"color,omitempty"`
	Rating            int    `json:"rating,omitempty"`
	HighestRatingType string `json:"highestRatingType,omitempty"`
}

type TopTrackMembers struct {
	Data []*RankItem `json:"data,omitempty"`
}

func (api *MembersApi) TopTrack(tt TrackType, opt *TopTrackListOptions) (*TopTrackMembers, *Response, error) {
	path, err := urlParameters(fmt.Sprintf("users/tops/%v", tt), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := api.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	topMembers := new(TopTrackMembers)
	resp, err := api.client.Do(req, topMembers)
	if err != nil {
		return nil, resp, err
	}

	return topMembers, resp, err
}
