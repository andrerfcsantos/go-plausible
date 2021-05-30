package plausible

import "strconv"

// BreakdownQuery represents an API query for detailed information about a property over a period of time.
// In an breakdown query, the Property field and the Period fields are mandatory, all the others are optional.
type BreakdownQuery struct {
	// Property is the property name for which the breakdown result will be about.
	// This field is mandatory.
	Property PropertyName
	// Period is the period of time to consider for the results.
	// This field is mandatory.
	Period TimePeriod
	// Metrics is a list of metrics for which to include in the results.
	// This field is optional.
	Metrics Metrics
	// Limit limits the number of results to be returned.
	// This field is optional.
	Limit int
	// Page indicates the page number for which to fetch the results. Page numbers start at 1.
	// This field is optional.
	Page int
	// Filters is a filter over properties to narrow down the breakdown results.
	// This field is optional.
	Filters Filter
}

// Validate tells whether the query is valid or not.
// If the query is not valid, a string explaining why the query is not valid will be returned.
func (bq *BreakdownQuery) Validate() (ok bool, invalidReason string) {
	if bq.Property.IsEmpty() {
		return false, "a property must be specified for a breakdown query"
	}
	if bq.Period.IsEmpty() {
		return false, "a period must be specified for a breakdown query"
	}

	return true, ""
}

func (bq *BreakdownQuery) toQueryArgs() QueryArgs {
	queryArgs := QueryArgs{}

	queryArgs.Merge(bq.Property.toQueryArgs())
	queryArgs.Merge(bq.Period.toQueryArgs())

	if !bq.Metrics.IsEmpty() {
		queryArgs.Merge(bq.Metrics.toQueryArgs())
	}

	if bq.Limit != 0 {
		queryArgs.Add(QueryArg{Name: "limit", Value: strconv.Itoa(bq.Limit)})
	}

	if bq.Page != 0 {
		queryArgs.Add(QueryArg{Name: "page", Value: strconv.Itoa(bq.Page)})
	}

	if !bq.Filters.IsEmpty() {
		queryArgs.Merge(bq.Filters.toQueryArgs())
	}

	return queryArgs
}

// PropertyResult contains the value of a property for an entry in the breakdown query results.
// At any moment, only the field corresponding to the property indicated in the query must be used.
// All the other fields will be empty.
type PropertyResult struct {
	// Name contains a value of the event name property.
	// This value must only be only if the breakdown query was for this property.
	Name string `json:"name"`
	// Page contains a value of the event page property.
	// This value must only be only if the breakdown query was for this property.
	Page string `json:"page"`
	// Page contains a value of the visit source property.
	// This value must only be only if the breakdown query was for this property.
	Source string `json:"source"`
	// Referrer contains a value of the visit referrer property.
	// This value must only be only if the breakdown query was for this property.
	Referrer string `json:"referrer"`
	// UtmMedium contains a value of the utm medium property.
	// This value must only be only if the breakdown query was for this property.
	UtmMedium string `json:"utm_medium"`
	// UtmSource contains a value of the utm source property.
	// This value must only be only if the breakdown query was for this property.
	UtmSource string `json:"utm_source"`
	// UtmCampaign contains a value of the utm campaign property.
	// This value must only be only if the breakdown query was for this property.
	UtmCampaign string `json:"utm_campaign"`
	// Device contains a value of the device property.
	// This value must only be only if the breakdown query was for this property.
	Device string `json:"device"`
	// Browser contains a value of the browser property.
	// This value must only be only if the breakdown query was for this property.
	Browser string `json:"browser"`
	// BrowserVersion contains a value of the browser version property.
	// This value must only be only if the breakdown query was for this property.
	BrowserVersion string `json:"browser_version"`
	// OS contains a value of the operating system property.
	// This value must only be only if the breakdown query was for this property.
	OS string `json:"os"`
	// OSVersion contains a value of the operating system version property.
	// This value must only be only if the breakdown query was for this property.
	OSVersion string `json:"os_version"`
	// Country contains a value of the country property.
	// This value must only be only if the breakdown query was for this property.
	Country string `json:"country"`
}

// BreakdownResultEntry represents an entry in a breakdown query result.
type BreakdownResultEntry struct {
	// PropertyResult contains the property value associated with this entry
	PropertyResult
	// MetricsResult contains the metric information associated with this entry
	MetricsResult
}

type rawBreakdownResponse struct {
	Results []BreakdownResultEntry `json:"results"`
}

// BreakdownResult represents a result
type BreakdownResult []BreakdownResultEntry
