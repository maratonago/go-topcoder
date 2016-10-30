package topcoder

type DataApi struct {
	client *Client
}

type SrmType string

const (
	Short      SrmType = "Single Round Match"
	Long       SrmType = "Long Round"
	Tournament SrmType = "Tournament"
)

// TODO - figure out Status meaning
type SrmStatus string

const (
	SrmStatusA SrmStatus = "a"
	SrmStatusF SrmStatus = "f"
	SrmStatusP SrmStatus = "p"
)

type SRMScheduleOptions struct {
	Statuses []SrmStatus `url:"statuses,omitempty"`
	Types    []SrmType   `url:"types,omitempty"`
	// TODO - add filters on time fields
	ListOptions
}

type ScheduledSrm struct {
	RoundId     uint      `json:"roundId,omitempty"`
	Name        string    `json:"name,omitempty"`
	ShorName    string    `json:"shortName,omitempty"`
	ContestName string    `json:"contestName,omitempty"`
	RoundType   SrmType   `json:"roundType,omitempty"`
	Status      SrmStatus `json:"status,omitempty"`
	// TODO - add time fields
}

type SrmSchedule struct {
	Data []*ScheduledSrm `json:"data,omitempty"`
}

func (api *DataApi) GetSrmSchedule(opt *SRMScheduleOptions) (*SrmSchedule, *Response, error) {
	path, err := urlParameters("data/srm/schedule", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := api.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	schedule := new(SrmSchedule)
	resp, err := api.client.Do(req, schedule)
	if err != nil {
		return nil, resp, err
	}

	return schedule, resp, err
}
