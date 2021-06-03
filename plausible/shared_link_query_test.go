package plausible_test

import (
	"os"
	"strings"
	"testing"

	"github.com/andrerfcsantos/go-plausible/plausible"
)

func TestSharedLinkQuery(t *testing.T) {
	tests := []struct {
		name       string
		query      plausible.SharedLinkQuery
		shouldFail bool
	}{
		{
			name: "shared link",
			query: plausible.SharedLinkQuery{
				Name: "go-plausible",
			},
			shouldFail: false,
		},
	}

	token := os.Getenv("PLAUSIBLE_PROVISIONING_TOKEN")
	rawDomains := os.Getenv("PLAUSIBLE_DOMAINS")

	if token == "" || rawDomains == "" {
		t.Skipf("no token or domain present in the environment variables")
	}

	siteStr := strings.Split(rawDomains, ",")[0]

	client := plausible.NewClient(token)
	site := client.Site(siteStr)

	for _, test := range tests {
		_, err := site.SharedLink(test.query)
		if err != nil && !test.shouldFail {
			t.Fatalf("test '%s' failed, got unexpected error: %v", test.name, err)
		}
	}
}
