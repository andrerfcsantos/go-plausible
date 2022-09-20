package plausible

// Metric represents a Plausible metric. Metrics are used as parts of queries to request specific information.
// Most users won't need to work with this type directly and can just use the metric constant values
// defined on this package when a Metric value is needed.
type Metric string

// Metric values:
const (
	// Visitors represents the number of visitors metric
	Visitors = Metric("visitors")
	// PageViews represents the number of page views metric
	PageViews = Metric("pageviews")
	// BounceRate represents the bounce rate metric
	BounceRate = Metric("bounce_rate")
	// VisitDuration represents the visit duration metric
	VisitDuration = Metric("visit_duration")
	// Visits represents the number of visits/sessions metric
	Visits = Metric("visits")
	// Events represents the number of events (pageviews + custom events) metric
	Events = Metric("events")
)

// AllMetrics is an utility function that returns all the metrics.
// This can be used to easily request information about all the metrics when building a query.
// However, please note that querying all metrics is not allowed in all type of queries.
func AllMetrics() Metrics {
	return Metrics{
		Visitors, PageViews, BounceRate, VisitDuration, Visits,
	}
}

// Metrics represents a list of metrics.
type Metrics []Metric

func (m *Metrics) toQueryVal() string {
	s := ""
	size := m.Count()

	for i := 0; i < size; i++ {
		s += string((*m)[i])
		if i != size-1 {
			s += ","
		}
	}
	return s
}

// Count returns the number of metrics in the list.
func (m *Metrics) Count() int {
	return len(*m)
}

// IsEmpty tells whether the list of metrics has any metrics.
func (m *Metrics) IsEmpty() bool {
	return m.Count() == 0
}

func (m *Metrics) toQueryArgs() QueryArgs {
	return QueryArgs{
		{Name: "metrics", Value: m.toQueryVal()},
	}
}
