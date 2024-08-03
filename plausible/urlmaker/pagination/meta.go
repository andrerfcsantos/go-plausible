package pagination

// Meta contains pagination information related with a response
type Meta struct {
	// After is the domain id that appears before all the records in the current page
	After string `json:"after"`
	// Before is the domain id that appears after all the records in the current page
	Before string `json:"before"`
	// Limit limits the number of records in a page
	Limit int `json:"limit"`
}
