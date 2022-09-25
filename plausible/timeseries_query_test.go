package plausible

import (
	"os"
	"strings"
	"testing"
)

func TestUnitValidateTimeseriesQuery(t *testing.T) {
	tests := []struct {
		name    string
		query   TimeseriesQuery
		isValid bool
	}{
		{
			name: "valid timeseries query",
			query: TimeseriesQuery{
				Period:   DayPeriod(),
				Filters:  NewFilter().ByVisitOs("Windows"),
				Metrics:  AllMetrics(),
				Interval: DateInterval,
			},
			isValid: true,
		},
		{
			name: "invalid timeseries due to missing time period",
			query: TimeseriesQuery{
				Filters:  NewFilter().ByVisitOs("Windows"),
				Metrics:  AllMetrics(),
				Interval: DateInterval,
			},
			isValid: false,
		},
	}

	for _, test := range tests {
		valid, _ := test.query.Validate()
		if valid && !test.isValid {
			t.Fatalf("test '%s' is valid, but was expected to fail", test.name)
		}
		if !valid && test.isValid {
			t.Fatalf("test '%s' is invalid, but was expected to succeed", test.name)
		}
	}

}

func TestUnitToQueryArgsTimeseriesQuery(t *testing.T) {
	tests := []struct {
		name              string
		query             TimeseriesQuery
		expectedQueryArgs QueryArgs
		isValid           bool
	}{
		{
			name: "valid breakdown query",
			query: TimeseriesQuery{
				Period:   DayPeriod(),
				Filters:  NewFilter().ByVisitOs("Windows"),
				Metrics:  AllMetrics(),
				Interval: DateInterval,
			},
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "period", Value: "day"},
				QueryArg{Name: "filters", Value: "visit:os==Windows"},
				QueryArg{Name: "metrics", Value: "visitors,pageviews,bounce_rate,visit_duration,visits"},
				QueryArg{Name: "interval", Value: "date"},
			},
			isValid: true,
		},
	}

	for _, test := range tests {
		got := test.query.toQueryArgs()
		if got.Count() != test.expectedQueryArgs.Count() {
			t.Fatalf("test '%s' failed because expected and actual query args have different sizes %d != %d: expected %#v got: %#v",
				test.name,
				got.Count(), test.expectedQueryArgs.Count(),
				got, test.expectedQueryArgs)
		}
		size := got.Count()
		for i := 0; i < size; i++ {
			if got[i].Name != test.expectedQueryArgs[i].Name {
				t.Fatalf("test '%s' failed because expected and actual query argument names are different at position %d: %s != %s",
					test.name, i, got[i].Name, test.expectedQueryArgs[i].Name)
			}

			if got[i].Value != test.expectedQueryArgs[i].Value {
				t.Fatalf("test '%s' failed because expected and actual query argument values for %s are different at position %d: %s != %s",
					test.name, got[i].Name, i, got[i].Value, test.expectedQueryArgs[i].Value)
			}
		}
	}

}

func TestIntegrationTimeseries(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		query      TimeseriesQuery
		shouldFail bool
	}{
		{
			name: "valid aggregate query with no period should succeed",
			query: TimeseriesQuery{
				Period:   MonthPeriod(),
				Filters:  NewFilter().ByVisitBrowser("Firefox"),
				Metrics:  AllMetrics(),
				Interval: DateInterval,
			},
			shouldFail: false,
		},
		{
			name: "invalid aggregate query with no period should fail",
			query: TimeseriesQuery{
				Filters:  NewFilter().ByVisitBrowser("Firefox"),
				Metrics:  AllMetrics(),
				Interval: DateInterval,
			},
			shouldFail: true,
		},
	}

	token := os.Getenv("PLAUSIBLE_TOKEN")
	rawDomains := os.Getenv("PLAUSIBLE_DOMAINS")

	if token == "" || rawDomains == "" {
		t.Skipf("no token or domain present in the environment variables")
	}

	siteStr := strings.Split(rawDomains, ",")[0]

	client := NewClient(token)
	site := client.Site(siteStr)

	for _, test := range tests {
		_, err := site.Timeseries(test.query)
		if err != nil && !test.shouldFail {
			t.Fatalf("test '%s' failed but was expected to suceed: %v", test.name, err)
		}

		if err == nil && test.shouldFail {
			t.Fatalf("test '%s' was expected to fail, but suceeded: %v", test.name, err)
		}
	}
}
