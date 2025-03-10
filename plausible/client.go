package plausible

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/andrerfcsantos/go-plausible/plausible/urlmaker/pagination"
	"github.com/valyala/fasthttp"
)

// DefaultBaseURL contains the default base url for the plausible API.
const DefaultBaseURL = "https://plausible.io/api/v1/"

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
		baseURL: DefaultBaseURL,
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

// BaseURL returns the base URL this client is using.
func (c *Client) BaseURL() string {
	return c.baseURL
}

// Token returns the token this client is using.
func (c *Client) Token() string {
	return c.token
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

// acquireRequest is a generic function for acquiring a new request that is already authenticated with the client token
// and has the base URL set to the default base URL.
func (c *Client) acquireRequest(method, endpoint string, queries QueryArgs, formData QueryArgs) (*fasthttp.Request, error) {
	return c.acquireRequestWithBaseURl(c.baseURL, method, endpoint, queries, formData)
}

// acquireRequest is the same as acquireRequest but allows the specification of a base URL.
// This can be used for endpoints that do not follow the
func (c *Client) acquireRequestWithBaseURl(baseURL, method, endpoint string, queries QueryArgs, formData QueryArgs) (*fasthttp.Request, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(baseURL + endpoint)
	req.Header.SetMethod(method)
	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("User-Agent", "go-plausible")

	for _, q := range queries {
		req.URI().QueryArgs().Add(q.Name, q.Value)
	}

	if formData != nil && formData.Count() > 0 {
		body := &bytes.Buffer{}
		mpwriter := multipart.NewWriter(body)

		for _, q := range formData {
			fw, err := mpwriter.CreateFormField(q.Name)
			if err != nil {
				fasthttp.ReleaseRequest(req)
				return nil, fmt.Errorf("creating form field '%s': %v", q.Name, err)
			}
			_, err = io.Copy(fw, strings.NewReader(q.Value))
			if err != nil {
				fasthttp.ReleaseRequest(req)
				return nil, fmt.Errorf("creating copying form field '%s' to form writer: %v", q.Name, err)
			}
		}

		err := mpwriter.Close()
		if err != nil {
			fasthttp.ReleaseRequest(req)
			return nil, fmt.Errorf("closing form multipart writer: %w", err)
		}

		req.Header.Set("Content-Type", mpwriter.FormDataContentType())
		req.SetBody(body.Bytes())
	}

	return req, nil
}

// CreateNewSite creates a new site in Plausible.
//
// Note: This endpoint requires an API token with permissions to use the sites provisioning API.
// Check https://plausible.io/docs/sites-api for more info
func (c *Client) CreateNewSite(siteRequest CreateSiteRequest) (CreateSiteResult, error) {
	ok, invalidReason := siteRequest.Validate()
	if !ok {
		return CreateSiteResult{}, errors.New("invalid request for new site: " + invalidReason)
	}
	req, err := c.acquireRequest("POST", "sites", nil, siteRequest.toFormArgs())
	if err != nil {
		return CreateSiteResult{}, fmt.Errorf("error acquiring request: %v", err)
	}

	data, err := doRequest(c.client, req)
	if err != nil {
		return CreateSiteResult{}, fmt.Errorf("error performing request to create new site: %v", err)
	}

	var res CreateSiteResult
	err = json.Unmarshal(data, &res)
	if err != nil {
		return CreateSiteResult{}, fmt.Errorf("error parsing create site response: %w", err)
	}

	return res, nil
}

// ListSites lists existing sites in Plausible
func (c *Client) ListSites(pagOptions ...pagination.Option) (ListSitesResult, error) {

	paginator := pagination.NewPaginator(pagOptions...)
	qArgs := QueryArgsFromPaginator(paginator)

	req, err := c.acquireRequest("GET", "sites", qArgs, nil)
	if err != nil {
		return ListSitesResult{}, fmt.Errorf("error acquiring request: %v", err)
	}

	data, err := doRequest(c.client, req)
	if err != nil {
		return ListSitesResult{}, fmt.Errorf("error performing request to list sites: %v", err)
	}

	var res ListSitesResult
	err = json.Unmarshal(data, &res)
	if err != nil {
		return ListSitesResult{}, fmt.Errorf("error parsing list sites response: %w", err)
	}

	return res, nil
}

// PushEvent records an event on plausible
func (c *Client) PushEvent(ev EventRequest) ([]byte, error) {
	req, err := c.acquireEventRequest(ev)
	if err != nil {
		return nil, fmt.Errorf("acquiring event request from client: %w", err)
	}
	return doRequest(c.client, req)
}
