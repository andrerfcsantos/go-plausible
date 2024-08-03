package plausible

import "github.com/andrerfcsantos/go-plausible/plausible/urlmaker/pagination"

// SiteResult contains the details for a Site
type SiteResult struct {
	// Domain of the site.
	Domain string `json:"domain"`
	// Timezone of the newly created site.
	Timezone string `json:"timezone"`
}

// ListSitesResult is the result of a request to list sites.
type ListSitesResult struct {
	// Sites is the list of sites in a response
	Sites []SiteResult `json:"sites"`
	// Meta is the pagination meta information of a page
	Meta pagination.Meta `json:"meta"`
}
