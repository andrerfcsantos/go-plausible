package plausible

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestUnitValidateAggregateQuery(t *testing.T) {
	tests := []struct {
		name    string
		query   AggregateQuery
		isValid bool
	}{
		{
			name: "valid aggregate aggregate query",
			query: AggregateQuery{
				Period:  DayPeriod(),
				Metrics: AllMetrics(),
			},
			isValid: true,
		},
		{
			name: "invalid aggregate query due to missing period",
			query: AggregateQuery{
				Metrics: AllMetrics(),
			},
			isValid: false,
		},
		{
			name: "invalid aggregate query due to missing metrics",
			query: AggregateQuery{
				Period: DayPeriod(),
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

func TestUnitToQueryArgsAggregateQuery(t *testing.T) {
	tests := []struct {
		name              string
		query             AggregateQuery
		expectedQueryArgs QueryArgs
		isValid           bool
	}{
		{
			name: "valid aggregate query",
			query: AggregateQuery{
				Period: CustomPeriod(
					Date{Day: 1, Month: 1, Year: 2021},
					Date{Day: 1, Month: 2, Year: 2021},
				),
				Metrics:               AllMetrics(),
				Filters:               NewFilter().ByVisitOs("Windows").ByEventPage("/"),
				ComparePreviousPeriod: true,
			},
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "period", Value: "custom"},
				QueryArg{Name: "date", Value: "2021-01-01,2021-02-01"},
				QueryArg{Name: "metrics", Value: "visitors,pageviews,bounce_rate,visit_duration,visits"},
				QueryArg{Name: "filters", Value: "visit:os==Windows;event:page==/"},
				QueryArg{Name: "compare", Value: "previous_period"},
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

func TestUnitToAggregateResultConversion(t *testing.T) {
	tests := []struct {
		name               string
		rawAggregateResult rawAggregateResult
		expectedConversion AggregateResult
		isValid            bool
	}{
		{
			name: "valid aggregate rawAggregateResult",
			rawAggregateResult: rawAggregateResult{
				Result: nil,
			},
			expectedConversion: AggregateResult{},
			isValid:            true,
		},
	}
	for _, test := range tests {
		aggResult := test.rawAggregateResult.toAggregateResult()
		if test.expectedConversion != aggResult {
			t.Fatalf("test '%s' failed %#v != %#v", test.name, test.expectedConversion, aggResult)
		}

	}

}

func TestIntegrationAggregateQuery(t *testing.T) {
	t.Parallel()
	year, month, day := time.Now().Date()

	tests := []struct {
		name       string
		query      AggregateQuery
		shouldFail bool
	}{
		{
			name: "invalid aggregate query with no period should fail",
			query: AggregateQuery{
				Metrics:               AllMetrics(),
				Filters:               NewFilter().ByVisitBrowser("Firefox"),
				ComparePreviousPeriod: false,
			},
			shouldFail: true,
		},
		{
			name: "invalid aggregate query with no metrics should fail",
			query: AggregateQuery{
				Period:                Last6Months(),
				Filters:               NewFilter().ByVisitBrowser("Firefox"),
				ComparePreviousPeriod: false,
			},
			shouldFail: true,
		},
		{
			name: "valid aggregate query that should succeed",
			query: AggregateQuery{
				Period: Last6Months().FromDate(Date{Day: day, Month: int(month), Year: year}),
				Metrics: Metrics{
					Visitors,
				},
			},
			shouldFail: false,
		},
		{
			name: "valid aggregate query with filters that should succeed",
			query: AggregateQuery{
				Period: Last7Days(),
				Metrics: Metrics{
					Visitors,
				},
				Filters: NewFilter().ByEventName("pageview"),
			},
			shouldFail: false,
		},
		{
			name: "valid aggregate query with comparison and with previous period that should succeed",
			query: AggregateQuery{
				Period: Last30Days(),
				Metrics: Metrics{
					Visitors,
				},
				ComparePreviousPeriod: true,
			},
			shouldFail: false,
		},
		{
			name: "valid aggregate query with comparison and with previous period that should succeed",
			query: AggregateQuery{
				Period: Last6Months().OfDate(Date{Day: day, Month: int(month), Year: year}),
				Metrics: Metrics{
					Visitors,
				},
				ComparePreviousPeriod: true,
			},
			shouldFail: false,
		},
		{
			name: "valid aggregate query with comparison and with previous period that should succeed",
			query: AggregateQuery{
				Period: Last12Months(),
				Metrics: Metrics{
					Visitors,
				},
			},
			shouldFail: false,
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
		_, err := site.Aggregate(test.query)
		if err != nil && !test.shouldFail {
			t.Fatalf("test '%s' failed but was expected to suceed: %v", test.name, err)
		}

		if err == nil && test.shouldFail {
			t.Fatalf("test '%s' was expected to fail, but suceeded: %v", test.name, err)
		}
	}
}
