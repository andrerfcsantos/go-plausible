package plausible

// TimeseriesQuery represents an API query for time series information over a period of time.
// In an aggregate query, the Metrics field is mandatory, all the others are optional.
type TimeseriesQuery struct {
	// Period to consider for the time series query.
	// The result will include results over this period of time.
	// This field is mandatory.
	Period TimePeriod
	// Filters is a filter over properties to narrow down the time series results.
	// This field is optional.
	Filters Filter
	// Metrics to be included in the time series information.
	// This field is optional.
	Metrics Metrics
	// Interval of time to consider for the time series result.
	// This field is optional.
	Interval TimeInterval
}

// Validate tells whether the query is valid or not.
// If the query is not valid, a string explaining why the query is not valid will be returned.
func (aq *TimeseriesQuery) Validate() (ok bool, invalidReason string) {

	if aq.Period.IsEmpty() {
		return false, "a period must be specified for a timeseries query"
	}

	return true, ""
}

func (aq *TimeseriesQuery) toQueryArgs() QueryArgs {
	queryArgs := QueryArgs{}

	queryArgs.Merge(aq.Period.toQueryArgs())

	if !aq.Filters.IsEmpty() {
		queryArgs.Merge(aq.Filters.toQueryArgs())
	}

	if !aq.Metrics.IsEmpty() {
		queryArgs.Merge(aq.Metrics.toQueryArgs())
	}

	if !aq.Interval.IsEmpty() {
		queryArgs.Merge(aq.Interval.toQueryArgs())
	}

	return queryArgs
}

// TimeseriesResult represents the result of a time series query.
type TimeseriesResult []TimeseriesDataPoint

type rawTimeseriesResponse struct {
	Results []TimeseriesDataPoint `json:"results"`
}

// MetricsResult contains the results for metrics data.
type MetricsResult struct {
	// BounceRateRaw contains information about the bounce rate.
	// This field must only be used if the query requested the bounce rate metric.
	// Even when the query requests information for the bounce rate metric, some data points can
	// have this field as nil.
	// If you don't care about the nil value, use the BounceRate function to get this value.
	BounceRateRaw *float64 `json:"bounce_rate"`

	// Pageviews contains information about the number of page views.
	// This field must only be used if the query requested the page views metric.
	Pageviews int `json:"pageviews"`

	// VisitDurationRaw contains information about the visit duration.
	// Only use this field if the query requested the visit duration metric.
	// Even when the query requests information for the visit duration metric, some data points can
	// have this field as nil.
	// If you don't care about the nil value, use the VisitDuration function to get this value.
	VisitDurationRaw *float64 `json:"visit_duration"`

	// Visitors contains information about the number of visitors.
	// This field must only be used if the query requested the visitors metric.
	Visitors int `json:"visitors"`

	// Visits contains information about the number of visits per session.
	// This field must only be used if the query requested the visits metric.
	Visits int `json:"visits"`
}

// BounceRate returns the bounce rate associated with this result.
// It will return 0 (zero) if the bounce rate information is not present.
func (mr *MetricsResult) BounceRate() float64 {
	if mr.BounceRateRaw == nil {
		return 0
	}
	return *mr.BounceRateRaw
}

// VisitDuration returns the visit duration associated with this result.
// It will return 0 (zero) if the visit duration information is not present.
func (mr *MetricsResult) VisitDuration() float64 {
	if mr.VisitDurationRaw == nil {
		return 0
	}
	return *mr.VisitDurationRaw
}

// TimeseriesDataPoint represents a data point in a time series result.
type TimeseriesDataPoint struct {
	// Date is a string containing information about the date this result refers to in the format of "yyyy-mm-dd".
	// For some queries, this string will also include information about an hour of day, in the format "yyyy-mm-dd hh:mm:ss"
	Date string `json:"date"`
	// MetricsResult contains the metric results for the metrics included in the query
	MetricsResult
}
