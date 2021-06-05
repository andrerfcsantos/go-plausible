package plausible

import "testing"

func TestUnitValidateSharedLinkRequest(t *testing.T) {
	tests := []struct {
		name    string
		request SharedLinkRequest
		isValid bool
	}{
		{
			name: "valid new shared link request that should succeed",
			request: SharedLinkRequest{
				Name: "Friends Link",
			},
			isValid: true,
		},
		{
			name:    "invalid new shared link without name that should fail",
			request: SharedLinkRequest{},
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

func TestUnitToFormArgsNewSharedLinkRequest(t *testing.T) {

	tests := []struct {
		name             string
		siteName         string
		request          SharedLinkRequest
		expectedFormArgs QueryArgs
	}{
		{
			name:     "valid new shared link request with name",
			siteName: "example.com",
			request: SharedLinkRequest{
				Name: "plausible-link",
			},
			expectedFormArgs: QueryArgs{
				QueryArg{Name: "site_id", Value: "example.com"},
				QueryArg{Name: "name", Value: "plausible-link"},
			},
		},
		{
			name:     "invalid new shared link request without name",
			request:  SharedLinkRequest{},
			siteName: "example.com",
			expectedFormArgs: QueryArgs{
				QueryArg{Name: "site_id", Value: "example.com"},
				QueryArg{Name: "name", Value: ""},
			},
		},
	}
	for _, test := range tests {
		actualFormArgs := test.request.toFormArgs(test.siteName)

		equal := actualFormArgs.equalTo(test.expectedFormArgs)
		if !equal {
			t.Fatalf("test '%s' failed: non-equal form args %v and %v",
				test.name, test.expectedFormArgs, actualFormArgs)
		}
	}

}
