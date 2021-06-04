package plausible

import (
	"os"
	"strings"
	"testing"
)

func TestIntegrationCurrentVisitors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		shouldFail bool
	}{
		{
			name:       "test",
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
		_, err := site.CurrentVisitors()
		if err != nil && !test.shouldFail {
			t.Fatalf("test '%s' failed but was expected to suceed: %v", test.name, err)
		}

		if err == nil && test.shouldFail {
			t.Fatalf("test '%s' was expected to fail, but suceeded: %v", test.name, err)
		}
	}
}
