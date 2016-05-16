package topcoder

import (
	"fmt"
)

type UserProfileApi struct {
	client *Client
}

type PublicProfile struct {
	Handle         *string  `json:"handle,omitempty"`
	Country        *string  `json:"country,omitempty"`
	MemberSince    *string  `json:"memberSince,omitempty"`
	Quote          *string  `json:"quote,omitempty"`
	PhotoLink      *string  `json:"photoLink,omitempty"`
	Copilot        *bool    `json:"copilot,omitempty"`
	OverallEarning *float64 `json:"overallEarning,omitempty"`
	RatingSummary  []struct {
		Name       *string `json:"name,omitempty"`
		Rating     *int    `json:"rating,omitempty"`
		ColorStyle *string `json:"colorStyle,omitempty"`
	} `json:"ratingSummary,omitempty"`
	Achievements []struct {
		Date        *string `json:"date,omitempty"`
		Description *string `json:"description,omitempty"`
	} `json:"Achievements,omitempty"`
}

func (api *UserProfileApi) PublicProfile(handle string) (*PublicProfile, *Response, error) {
	path := fmt.Sprintf("users/%v", handle)

	req, err := api.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	profile := new(PublicProfile)
	resp, err := api.client.Do(req, profile)
	if err != nil {
		return nil, resp, err
	}

	return profile, resp, err
}
