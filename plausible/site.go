package plausible

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

// Site represents a site added to plausible and implements a client
// for all stats requests related with the site.
//
// Site is safe for concurrent use.
type Site struct {
	token           string
	id              string
	httpClient      *fasthttp.Client
	plausibleClient *Client
}

// ID returns the ID of the site.
func (s *Site) ID() string {
	return s.id
}

func (s *Site) acquireRequest(method, endpoint string, queries QueryArgs) *fasthttp.Request {
	req := s.plausibleClient.acquireRequest(method, endpoint, queries)
	req.URI().QueryArgs().Add("site_id", s.id)
	return req
}

func (s *Site) doRequest(method, endpoint string, queries QueryArgs) ([]byte, error) {
	req := s.acquireRequest(method, endpoint, queries)
	data, err := doRequest(s.httpClient, req)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CurrentVisitors gets the current visitors for the site.
func (s *Site) CurrentVisitors() (int, error) {
	data, err := s.doRequest("GET", "stats/realtime/visitors", nil)
	if err != nil {
		return 0, fmt.Errorf("error performing current visitors request: %w", err)
	}

	visitors, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}

	return visitors, nil
}

// Aggregate performs an aggregate query.
// An aggregate query reports data for metrics aggregated over a period of time,
// eg, "total number of visitors/pageviews for a particular day".
func (s *Site) Aggregate(query AggregateQuery) (AggregateResult, error) {

	ok, invalidReason := query.Validate()
	if !ok {
		return AggregateResult{}, errors.New("invalid aggregate query: " + invalidReason)
	}

	data, err := s.doRequest("GET", "stats/aggregate", query.toQueryArgs())
	if err != nil {
		return AggregateResult{}, fmt.Errorf("error performing aggregate request: %w", err)
	}

	var res rawAggregateResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return AggregateResult{}, fmt.Errorf("error parsing aggregate response: %w", err)
	}

	return res.toAggregateResponse(), nil
}

// Timeseries performs a time series query.
// A time series query reports a list of data points over a period of time,
// where each data point contains data about metrics for that period of time.
// e.g, "total number of visitors and page views for each day in the last month".
func (s *Site) Timeseries(query TimeseriesQuery) (TimeseriesResult, error) {

	ok, invalidReason := query.Validate()
	if !ok {
		return TimeseriesResult{}, errors.New("invalid timeline query: " + invalidReason)
	}

	data, err := s.doRequest("GET", "stats/timeseries", query.toQueryArgs())
	if err != nil {
		return TimeseriesResult{}, fmt.Errorf("error performing timeline request: %w", err)
	}

	var res rawTimeseriesResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return TimeseriesResult{}, fmt.Errorf("error parsing timeline response: %w", err)
	}

	return res.Results, nil
}

// Breakdown performs a breakdown query.
// A breakdown query reports stats for the value of a given property over a period of time,
// e.g, "total number of visitors and page views for each operating system in the last month".
func (s *Site) Breakdown(query BreakdownQuery) (BreakdownResult, error) {

	ok, invalidReason := query.Validate()
	if !ok {
		return BreakdownResult{}, errors.New("invalid breakdown query: " + invalidReason)
	}

	data, err := s.doRequest("GET", "stats/breakdown", query.toQueryArgs())
	if err != nil {
		return BreakdownResult{}, fmt.Errorf("error performing breakdown request: %w", err)
	}

	var res rawBreakdownResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return BreakdownResult{}, fmt.Errorf("error parsing breakdown response: %w", err)
	}

	return res.Results, nil
}
