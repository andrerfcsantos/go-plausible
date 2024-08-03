package pagination

// Paginator holds information to paginate requests
type Paginator struct {
	// After is the domain id that appears before all the records in the desired page
	After string
	// Before is the domain id that appears after all the records in the desired page
	Before string
	// Limit sets the maximum records in the desired page
	Limit int
}

// NewPaginator creates a new paginator with the given options
func NewPaginator(options ...Option) *Paginator {
	paginator := &Paginator{}
	for _, opt := range options {
		opt(paginator)
	}
	return paginator
}
