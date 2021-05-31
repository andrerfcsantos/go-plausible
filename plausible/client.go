package plausible

import (
	"strings"

	"github.com/valyala/fasthttp"
)

// Client handles the interaction with the plausible API.
//
// The client must be initialized with a token using either NewClient or NewClientWithBaseURL.
// It's safe to use this client concurrently.
type Client struct {
	baseURL string
	token   string
	client  *fasthttp.Client
}

// NewClient returns a new API client with the given token.
// Calling this function is the way most users want to use to create and initialize a new client.
// This function does not make any network requests.
//
// This client will use the API located at https://plausible.io/api/v1/.
// If you need to use another base URL for the API, create a client using NewClientWithBaseURL instead.
func NewClient(token string) *Client {
	return &Client{
		baseURL: "https://plausible.io/api/v1/",
		token:   token,
		client:  &fasthttp.Client{},
	}
}

// NewClientWithBaseURL creates a new API token with a given token, similarly to NewClient,
// but also allows the specification of a base URL for the API.
// This function does not make any network requests.
//
// This allows the specification of an URL for a self-hosted API or another version of the API.
// The url must be a complete url as it must contain a schema, the domain for the API and the prefix path of the
// API, e.g. "https://plausible.io/api/v1/". Including a trailing / in the URL is optional.
func NewClientWithBaseURL(token string, baseURL string) *Client {
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	return &Client{
		baseURL: baseURL,
		token:   token,
		client:  &fasthttp.Client{},
	}
}

// Site returns a site handler for a given site ID. The returned handler can be used to query the API for
// information and statistics about the site. This function does not make any network requests.
func (c *Client) Site(siteID string) *Site {
	return &Site{
		token:           c.token,
		id:              siteID,
		httpClient:      c.client,
		plausibleClient: c,
	}
}

// acquireRequest returns a new request with the base fields for any request set up.
// All functions making requests should always call this function first.
func (c *Client) acquireRequest(method, endpoint string, queries QueryArgs) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(c.baseURL + endpoint)
	req.Header.SetMethod(method)
	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("User-Agent", "go-plausible")

	for _, q := range queries {
		req.URI().QueryArgs().Add(q.Name, q.Value)
	}

	return req
}
