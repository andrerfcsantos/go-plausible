package plausible

import "testing"

func TestFilterString(t *testing.T) {

	tests := []struct {
		filter   *Filter
		expected string
	}{
		{filter: NewFilter(), expected: ""},
		{filter: NewFilter().ByVisitBrowser("Firefox"), expected: "visit:browser==Firefox"},
		{filter: NewFilter().ByVisitBrowser("Firefox").ByVisitCountry("Portugal"), expected: "visit:browser==Firefox;visit:country==Portugal"},
	}

	for _, test := range tests {
		got := test.filter.toFilterString()
		if got != test.expected {
			t.Fatalf("TestFilterString: expected %s, got %s", test.expected, got)
		}
	}

}
