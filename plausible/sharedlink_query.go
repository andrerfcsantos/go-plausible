package plausible

// SharedLinkRequest represents a shared link request
type SharedLinkRequest struct {
	// Name is the name of the shared link.
	// This field is required.
	Name string
}

// Validate validates if a shared link request is valid
func (q *SharedLinkRequest) Validate() (bool, string) {
	if q.Name == "" {
		return false, "a link name must be specified for an shared link request"
	}

	return true, ""
}

func (q *SharedLinkRequest) toFormArgs(siteName string) QueryArgs {
	return QueryArgs{
		{Name: "site_id", Value: siteName},
		{Name: "name", Value: q.Name},
	}
}

// SharedLinkResult represents the result of a shared link request
type SharedLinkResult struct {
	// Name is the name of the shared link
	Name string `json:"name"`
	// URL is the URL for the shared link
	URL string `json:"url"`
}
