package plausible

// CreateSiteRequest represents a request to create a new site in Plausible.
type CreateSiteRequest struct {
	// Domain of the site to create.
	// This field is mandatory.
	Domain string
	// Timezone name according to the IANA database (e.g "Europe/London").
	// This field is optional and will default to "Etc/UTC".
	Timezone string
}

func (csr *CreateSiteRequest) toFormArgs() QueryArgs {
	res := QueryArgs{
		{Name: "domain", Value: csr.Domain},
	}

	if csr.Timezone != "" {
		res = append(res, QueryArg{Name: "timezone", Value: csr.Timezone})
	}

	return res
}

// Validate tells whether the request is valid or not.
// If the request is not valid, a string explaining why the request is not valid will be returned.
func (csr *CreateSiteRequest) Validate() (bool, string) {
	if csr.Domain == "" {
		return false, "a domain must be specified in a request to create a new site"
	}

	return true, ""
}

// CreateSiteResult is the result of a request to create a new site.
type CreateSiteResult struct {
	// Domain of the created site.
	Domain string `json:"domain"`
	// Timezone of the newly created site.
	Timezone string `json:"timezone"`
}
