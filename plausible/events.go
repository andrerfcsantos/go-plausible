package plausible

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"regexp"
)

var (
	apiVersionRegex *regexp.Regexp
)

func init() {
	apiVersionRegex = regexp.MustCompile(`/api/v\d/`)
}

// EventData represents the data of an event
type EventData struct {
	// Domain of the site
	Domain string `json:"domain"`
	// Name is the name of the event, e.g. 'pageview', otherwise a custom event
	Name string `json:"name"`

	// URL of the current request
	URL string `json:"url"`

	// Referrer to be associated with the event
	Referrer string `json:"referrer,omitempty"`
	// Props of the event
	Props map[string]string `json:"props,omitempty"`
	// Revenue associated with the event
	Revenue Revenue `json:"revenue,omitempty"`
}

// EventRequest represents the request for an event
type EventRequest struct {
	EventData
	// User Agent to be associated with request event.
	// This field is mandatory.
	UserAgent string
	// X-Forwarded-For header to be sent with the event. This field is optional.
	XForwardedFor string
	// AdditionalHeaders are additional headers to be included in the request. This field is optional.
	AdditionalHeaders map[string]string
	// IsDebuggingRequest tells if this is a debugging request. This field is optional.
	IsDebuggingRequest bool
}

// Revenue represents the revenue associated with an event
type Revenue struct {
	// Currency is
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

func (c *Client) acquireEventRequest(request EventRequest) (*fasthttp.Request, error) {
	if request.UserAgent == "" {

		return nil, fmt.Errorf("missing user agent information for the event request")
	}

	newBaseURL := apiVersionRegex.ReplaceAllString(c.baseURL, "/")
	req, err := c.acquireRequestWithBaseURl(newBaseURL, "POST", "api/event", nil, nil)
	if err != nil {

		return nil, fmt.Errorf("acquiring request from client for /api/event: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", request.UserAgent)

	if request.IsDebuggingRequest {
		req.Header.Add("X-Debug-Request", "true")
	}

	if request.XForwardedFor != "" {
		req.Header.Add("X-Forwarded-For", request.XForwardedFor)
	}

	for header, value := range request.AdditionalHeaders {
		req.Header.Add(header, value)
	}

	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(request.EventData)
	if err != nil {
		return nil, fmt.Errorf("encoding request body for event request: %w", err)
	}

	req.SetBody(body.Bytes())
	return req, nil
}
