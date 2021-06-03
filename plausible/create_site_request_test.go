package plausible_test

import (
	"os"
	"testing"

	"github.com/andrerfcsantos/go-plausible/plausible"
)

func TestCreateNewSiteRequestQuery(t *testing.T) {

	newDomain := os.Getenv("PLAUSIBLE_NEW_DOMAIN")

	if newDomain == "" {
		t.Skipf("no new domain is present in the environment variables")
	}

	tests := []struct {
		name       string
		query      plausible.CreateSiteRequest
		shouldFail bool
	}{
		{
			name: "create new site",
			query: plausible.CreateSiteRequest{
				Domain: newDomain,
			},
			shouldFail: false,
		},
	}

	token := os.Getenv("PLAUSIBLE_PROVISIONING_TOKEN")

	if token == "" {
		t.Skipf("no provisioning token present in the environment variables")
	}

	client := plausible.NewClient(token)

	for _, test := range tests {
		_, err := client.CreateNewSite(test.query)
		if err != nil && !test.shouldFail {
			t.Fatalf("test '%s' failed, got unexpected error: %v", test.name, err)
		}
	}
}
