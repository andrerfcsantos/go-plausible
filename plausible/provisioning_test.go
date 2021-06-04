// +build provisioning

package plausible

import (
	"os"
	"strings"
	"testing"
)

func TestIntegrationCreateNewSiteRequestQuery(t *testing.T) {
	t.Parallel()
	newDomain := os.Getenv("PLAUSIBLE_NEW_DOMAIN")

	if newDomain == "" {
		t.Skipf("no new domain is present in the environment variables")
	}

	tests := []struct {
		name       string
		query      CreateSiteRequest
		shouldFail bool
	}{
		{
			name:       "invalid new site request due to missing domain that should fail",
			query:      CreateSiteRequest{},
			shouldFail: true,
		},
		{
			name: "valid new site request that should succeed",
			query: CreateSiteRequest{
				Domain: newDomain,
			},
			shouldFail: false,
		},
		{
			name: "invalid new site request for existing site that should fail",
			query: CreateSiteRequest{
				Domain: newDomain,
			},
			shouldFail: true,
		},
	}

	token := os.Getenv("PLAUSIBLE_PROVISIONING_TOKEN")

	if token == "" {
		t.Skipf("no provisioning token present in the environment variables")
	}

	client := NewClient(token)

	for _, test := range tests {
		_, err := client.CreateNewSite(test.query)
		if err != nil && !test.shouldFail {
			t.Fatalf("test '%s' failed but was expected to suceed: %v", test.name, err)
		}

		if err == nil && test.shouldFail {
			t.Fatalf("test '%s' was expected to fail, but suceeded: %v", test.name, err)
		}
	}
}

func TestIntegrationSharedLinkRequest(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		query      SharedLinkRequest
		shouldFail bool
	}{
		{
			name:       "invalid new shared link request missing name that should fail",
			query:      SharedLinkRequest{},
			shouldFail: true,
		},
		{
			name: "valid bew shared link request that should succeed",
			query: SharedLinkRequest{
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

	client := NewClient(token)
	site := client.Site(siteStr)

	for _, test := range tests {
		_, err := site.SharedLink(test.query)
		if err != nil && !test.shouldFail {
			t.Fatalf("test '%s' failed, got unexpected error: %v", test.name, err)
		}
	}
}
