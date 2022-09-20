package plausible

import (
	"os"
	"strings"
	"testing"
)

func TestUnitValidateBreakdownQuery(t *testing.T) {
	tests := []struct {
		name    string
		query   BreakdownQuery
		isValid bool
	}{
		{
			name: "valid breakdown query",
			query: BreakdownQuery{
				Property: VisitBrowser,
				Period:   DayPeriod(),
				Metrics:  AllMetrics(),
				Limit:    1,
				Page:     1,
				Filters:  NewFilter().ByVisitOs("Windows"),
			},
			isValid: true,
		},
		{
			name: "invalid breakdown query due to missing property",
			query: BreakdownQuery{
				Metrics: AllMetrics(),
				Period:  DayPeriod(),
			},
			isValid: false,
		},
		{
			name: "invalid breakdown query due to missing time period",
			query: BreakdownQuery{
				Property: VisitBrowser,
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

func TestUnitToQueryArgsBreakdownQuery(t *testing.T) {
	tests := []struct {
		name              string
		query             BreakdownQuery
		expectedQueryArgs QueryArgs
		isValid           bool
	}{
		{
			name: "valid breakdown query",
			query: BreakdownQuery{
				Property: VisitBrowser,
				Period:   DayPeriod(),
				Metrics:  AllMetrics(),
				Limit:    1,
				Page:     1,
				Filters:  NewFilter().ByVisitOs("Windows").ByEventPage("/"),
			},
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "property", Value: "visit:browser"},
				QueryArg{Name: "period", Value: "day"},
				QueryArg{Name: "metrics", Value: "visitors,pageviews,bounce_rate,visit_duration,visits"},
				QueryArg{Name: "limit", Value: "1"},
				QueryArg{Name: "page", Value: "1"},
				QueryArg{Name: "filters", Value: "visit:os==Windows;event:page==/"},
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

func TestIntegrationBreakdownQuery(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		query      BreakdownQuery
		shouldFail bool
	}{
		{
			name: "valid breakdown query that should succeed",
			query: BreakdownQuery{
				Property: VisitOs,
				Period:   DayPeriod(),
			},
			shouldFail: false,
		},
		{
			name: "invalid breakdown query due to missing property should fail",
			query: BreakdownQuery{
				Period: DayPeriod(),
			},
			shouldFail: true,
		},
		{
			name: "invalid breakdown query due to missing period should fail",
			query: BreakdownQuery{
				Property: VisitOs,
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
		_, err := site.Breakdown(test.query)
		if err != nil && !test.shouldFail {
			t.Fatalf("test '%s' failed but was expected to suceed: %v", test.name, err)
		}

		if err == nil && test.shouldFail {
			t.Fatalf("test '%s' was expected to fail, but suceeded: %v", test.name, err)
		}
	}
}
