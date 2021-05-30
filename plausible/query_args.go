package plausible

// QueryArgs represents a list of query arguments.
type QueryArgs []QueryArg

// Add adds a query argument to the list.
func (qa *QueryArgs) Add(q QueryArg) *QueryArgs {
	*qa = append(*qa, q)
	return qa
}

// Merge adds together the query arguments in other lists to the current one.
func (qa *QueryArgs) Merge(qArgsList ...QueryArgs) *QueryArgs {
	for _, qArgs := range qArgsList {
		*qa = append(*qa, qArgs...)
	}
	return qa
}

// Count returns the number of query arguments in the list
func (qa *QueryArgs) Count() int {
	return len(*qa)
}

// QueryArg represents a query argument
type QueryArg struct {
	// Name is the name of the query argument
	Name string
	// Value is the value for the query argument
	Value string
}
