package topcoder

type SortOrderType string

const (
	Ascending  SortOrderType = "asc"
	Descending SortOrderType = "desc"
)

type ListOptions struct {
	PageIndex  int           `url:"pageIndex,omitempty"`
	PageSize   int           `url:"pageSize,omitempty"`
	SortColumn string        `url:"sortColumn,omitempty"`
	SortOrder  SortOrderType `url:"sortOrder,omitempty"`
}
