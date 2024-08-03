package pagination

// Option represents a pagination option.
// It is not meant to be created directly - instead, you should get
// an Option via the helper functions, like After, Before and Limit.
type Option func(*Paginator)

// After sets the 'after' pagination option
func After(after string) Option {
	return func(p *Paginator) {
		p.After = after
	}
}

// Before sets the 'before' pagination option
func Before(before string) Option {
	return func(p *Paginator) {
		p.Before = before
	}
}

// Limit sets the 'limit' pagination option
func Limit(limit int) Option {
	return func(p *Paginator) {
		p.Limit = limit
	}
}
