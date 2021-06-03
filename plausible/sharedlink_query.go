package plausible

// SharedLinkQuery represents a shared link query
type SharedLinkQuery struct {
	// Name is the name of the shared link.
	// This field is required.
	Name string
}

// Validate validates if a shared link query is valid
func (q *SharedLinkQuery) Validate() (bool, string) {
	if q.Name == "" {
		return false, "a link name must be specified for an shared link query"
	}

	return true, ""
}

func (q *SharedLinkQuery) toFormArgs(siteName string) QueryArgs {
	return QueryArgs{
		{Name: "site_id", Value: siteName},
		{Name: "name", Value: q.Name},
	}
}

// SharedLinkResult represents the result of a shared link query
type SharedLinkResult struct {
	// Name is the name of the shared link
	Name string
	// URL is the URL for the shared link
	URL string
}
