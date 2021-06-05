package plausible

import (
	"testing"
)

func TestUnitValidateNewSiteRequest(t *testing.T) {
	tests := []struct {
		name    string
		request CreateSiteRequest
		isValid bool
	}{
		{
			name: "valid new site request with domain and timezone that should succeed",
			request: CreateSiteRequest{
				Domain:   "mydomain.com",
				Timezone: "Europe/Lisbon",
			},
			isValid: true,
		},
		{
			name: "valid new site request with domain and timezone that should succeed",
			request: CreateSiteRequest{
				Domain: "mydomain.com",
			},
			isValid: true,
		},
		{
			name:    "invalid new site request without domain that should fail",
			request: CreateSiteRequest{},
			isValid: false,
		},
	}

	for _, test := range tests {
		valid, _ := test.request.Validate()
		if valid && !test.isValid {
			t.Fatalf("test '%s' is valid, but was expected to fail", test.name)
		}
		if !valid && test.isValid {
			t.Fatalf("test '%s' is invalid, but was expected to succeed", test.name)
		}
	}

}

func TestUnitToFormArgsNewSiteRequest(t *testing.T) {
	tests := []struct {
		name             string
		request          CreateSiteRequest
		expectedFormArgs QueryArgs
	}{
		{
			name: "valid new site request with domain and timezone",
			request: CreateSiteRequest{
				Domain:   "mydomain.com",
				Timezone: "Europe/Lisbon",
			},
			expectedFormArgs: QueryArgs{
				QueryArg{Name: "domain", Value: "mydomain.com"},
				QueryArg{Name: "timezone", Value: "Europe/Lisbon"},
			},
		},
		{
			name: "valid new site request with domain and without timezone",
			request: CreateSiteRequest{
				Domain: "mydomain.com",
			},
			expectedFormArgs: QueryArgs{
				QueryArg{Name: "domain", Value: "mydomain.com"},
			},
		},
	}
	for _, test := range tests {
		actualFormArgs := test.request.toFormArgs()
		equal := actualFormArgs.equalTo(test.expectedFormArgs)
		if !equal {
			t.Fatalf("test '%s' failed: non-equal form args %v and %v",
				test.name, test.expectedFormArgs, actualFormArgs)
		}
	}

}
