package plausible

import (
	"os"
	"strings"
	"testing"
)

func TestAggregateQuery(t *testing.T) {
	tests := []struct {
		name       string
		query      AggregateQuery
		shouldFail bool
	}{
		{
			name: "aggregate query with no period",
			query: AggregateQuery{
				Metrics:               AllMetrics(),
				Filters:               NewFilter().ByVisitBrowser("Firefox"),
				ComparePreviousPeriod: false,
			},
			shouldFail: true,
		},
		{
			name: "aggregate query with no metrics",
			query: AggregateQuery{
				Period:                DayPeriod(),
				Filters:               NewFilter().ByVisitBrowser("Firefox"),
				ComparePreviousPeriod: false,
			},
			shouldFail: true,
		},
		{
			name: "regular query that should work",
			query: AggregateQuery{
				Period: DayPeriod(),
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
			t.Fatalf("test '%s' failed, got unexpected error: %v", test.name, err)
		}
	}
}
