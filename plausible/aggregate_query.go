package plausible

// AggregateQuery represents an API query for aggregate information over a period of time for a given list of metrics.
// In an aggregate query, the Period field and the Metrics field are mandatory, all the others are optional.
type AggregateQuery struct {
	// Period to consider for the aggregate query. The result will include results over this period of time.
	// This field is mandatory.
	Period TimePeriod
	// Metrics to be included in the aggregation result.
	// This field is mandatory.
	Metrics Metrics
	// Filters is a filter over properties to narrow down the aggregation results.
	// This field is optional.
	Filters Filter
	// ComparePreviousPeriod tells whether to include a comparison with the previous period in the query result.
	// This field is optional and will default to false.
	ComparePreviousPeriod bool
}

// Validate tells whether the query is valid or not.
// If the query is not valid, a string explaining why the query is not valid will be returned.
func (aq *AggregateQuery) Validate() (ok bool, invalidReason string) {

	if aq.Period.IsEmpty() {
		return false, "a period must be specified for an aggregate query"
	}

	if aq.Metrics.IsEmpty() {
		return false, "at least one metric must be specified for an aggregate query"
	}

	return true, ""
}

func (aq *AggregateQuery) toQueryArgs() QueryArgs {
	queryArgs := QueryArgs{}

	queryArgs.Merge(aq.Period.toQueryArgs())
	queryArgs.Merge(aq.Metrics.toQueryArgs())
	if !aq.Filters.IsEmpty() {
		queryArgs.Merge(aq.Filters.toQueryArgs())
	}

	if aq.ComparePreviousPeriod {
		queryArgs.Add(QueryArg{Name: "compare", Value: "previous_period"})
	}

	return queryArgs
}

// AggregateResult represents the result of an aggregate query.
type AggregateResult struct {
	// BounceRate represents the bounce rate result for the query.
	// Only use this field if you included the BounceRate metric in your query.
	BounceRate float64 `json:"bounce_rate"`
	// BounceRateChange represents the bounce rate change compared to the previous period.
	// Only use this field if you included the BounceRate metric in your query and ComparePreviousPeriod was set to true.
	BounceRateChange float64 `json:"bounce_rate_change"`

	// Pageviews represents the page view result for the query.
	// Only use this field if you included the PageViews metric in your query.
	Pageviews int `json:"pageviews"`
	// PageviewsChange represents change in the number of pageviews compared to the previous period.
	// Only use this field if you included the PageViews metric in your query and ComparePreviousPeriod was set to true.
	PageviewsChange int `json:"pageviews_change"`

	// VisitDuration represents the visit duration result for the query.
	// Only use this field if you included the VisitDuration metric in your query
	VisitDuration float64 `json:"visit_duration"`
	// VisitDurationChange represents the visit duration change compared to the previous period.
	// Only use this field if you included the VisitVisitDuration metric in your query and ComparePreviousPeriod was set to true.
	VisitDurationChange float64 `json:"visit_duration_change"`

	// Visitors represents the number of visitors result for the query.
	// Only use this field if you included the Visitors metric in your query.
	Visitors int `json:"visitors"`
	// VisitorsChange represents the change in the number of visitors compared to the previous period.
	// Only use this field if you included the Visitors metric in your query and ComparePreviousPeriod was set to true.
	VisitorsChange int `json:"visitors_change"`

	// Visits represents the visits result for the query.
	// Only use this field if you included the Visits metric in your query
	Visits int `json:"visits"`
	// VisitsChange represents the visits change compared to the previous period.
	// Only use this field if you included the Visits metric in your query and ComparePreviousPeriod was set to true.
	VisitsChange int `json:"visits_change"`

	// Events represents the events result for the query.
	// Only use this field if you included the Events metric in your query
	Events int `json:"events"`
	// EventsChange represents the events change compared to the previous period.
	// Only use this field if you included the Events metric in your query and ComparePreviousPeriod was set to true.
	EventsChange int `json:"events_change"`
}

type rawAggregateResult struct {
	Result *struct {
		BounceRate struct {
			Change float64 `json:"change"`
			Value  float64 `json:"value"`
		} `json:"bounce_rate,omitempty"`
		Events struct {
			Change int `json:"change"`
			Value  int `json:"value"`
		} `json:"events,omitempty"`
		Pageviews struct {
			Change int `json:"change"`
			Value  int `json:"value"`
		} `json:"pageviews,omitempty"`
		VisitDuration struct {
			Change float64 `json:"change"`
			Value  float64 `json:"value"`
		} `json:"visit_duration,omitempty"`
		Visitors struct {
			Change int `json:"change"`
			Value  int `json:"value"`
		} `json:"visitors,omitempty"`
		Visits struct {
			Change int `json:"change"`
			Value  int `json:"value"`
		} `json:"visits,omitempty"`
	} `json:"results,omitempty"`
}

func (r *rawAggregateResult) toAggregateResult() AggregateResult {
	var res AggregateResult

	if r.Result == nil {
		return res
	}

	res.BounceRate = r.Result.BounceRate.Value
	res.BounceRateChange = r.Result.BounceRate.Change

	res.Events = r.Result.Events.Value
	res.EventsChange = r.Result.Events.Change

	res.Pageviews = r.Result.Pageviews.Value
	res.PageviewsChange = r.Result.Pageviews.Change

	res.VisitDuration = r.Result.VisitDuration.Value
	res.VisitDurationChange = r.Result.VisitDuration.Change

	res.Visitors = r.Result.Visitors.Value
	res.VisitorsChange = r.Result.Visitors.Change

	res.Visits = r.Result.Visits.Value
	res.VisitsChange = r.Result.Visits.Change

	return res
}
