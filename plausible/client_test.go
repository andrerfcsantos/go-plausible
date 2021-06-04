package plausible

import "testing"

func TestUnitClientCreation(t *testing.T) {
	tests := []struct {
		name            string
		client          *Client
		expectedToken   string
		expectedBaseURL string
	}{
		{
			name:            "valid client with default base url",
			client:          NewClient("a"),
			expectedToken:   "a",
			expectedBaseURL: DefaultBaseURL,
		},
		{
			name:            "valid client with custom base URL with trailing slash",
			client:          NewClientWithBaseURL("a", "https://mydomain.com/api/v1/"),
			expectedToken:   "a",
			expectedBaseURL: "https://mydomain.com/api/v1/",
		},
		{
			name:            "valid client with custom base URL without trailing slash",
			client:          NewClientWithBaseURL("a", "https://mydomain.com/api/v1"),
			expectedToken:   "a",
			expectedBaseURL: "https://mydomain.com/api/v1/",
		},
	}

	for _, test := range tests {
		token, url := test.client.Token(), test.client.BaseURL()
		if token != test.expectedToken {
			t.Fatalf("test '%s' failed: expected token: %s, got: %s", test.name, test.expectedToken, token)
		}
		if url != test.expectedBaseURL {
			t.Fatalf("test '%s' failed: expected base url: %s, got: %s", test.name, test.expectedToken, token)
		}
	}
}
