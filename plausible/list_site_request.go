package plausible

import "strconv"

type PaginationMeta struct {
	After  string `json:"after"`
	Before string `json:"before"`
	Limit  int    `json:"limit"`
}

type ListSitesRequest struct {
	After  string
	Before string
	Limit  int
}

func (lsr *ListSitesRequest) toQueryArgs() QueryArgs {
	queryArgs := QueryArgs{}

	if lsr.After != "" {
		queryArgs.Add(QueryArg{Name: "after", Value: lsr.After})
	}
	if lsr.Before != "" {
		queryArgs.Add(QueryArg{Name: "before", Value: lsr.Before})
	}
	if lsr.Limit != 0 {
		queryArgs.Add(QueryArg{Name: "limit", Value: strconv.Itoa(lsr.Limit)})
	}

	return queryArgs
}

type SiteResult struct {
	// Domain of the site.
	Domain string `json:"domain"`
	// Timezone of the newly created site.
	Timezone string `json:"timezone"`
}

// ListSitesResult is the result of a request to list sites.
type ListSitesResult struct {
	Sites []SiteResult   `json:"sites"`
	Meta  PaginationMeta `json:"meta"`
}
