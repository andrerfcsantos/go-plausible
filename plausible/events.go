package plausible

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/valyala/fasthttp"
)

type event struct {
	plausibleClient *Client
}

// EventRequest represents data to record events on plausible
type EventRequest struct {
	// the domain of your site (must match exactly)
	Domain string `json:"domain"`

	// the name of the event, e.g. 'pageview', otherwise a custom event
	Name string `json:"name"`

	// the URL of the current request
	URL *url.URL `json:"url"`

	// optional parameters
	Referrer string            `json:"referrer,omitempty"`
	Props    map[string]string `json:"props,omitempty"`
	Revenue  Revenue           `json:"revenue,omitempty"`

	// these are not passed on
	Headers map[string]string `json:"-"`
}

// Revenue is an optional structure to set currency and amount
type Revenue struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

func (e *event) acquireRequest(data EventRequest) (*fasthttp.Request, error) {
	userAgent, ok := data.Headers["user-agent"]
	if !ok {
		return nil, fmt.Errorf("missing user-agent header")
	}

	req, err := e.plausibleClient.acquireRequest("POST", "/api/event", QueryArgs{}, QueryArgs{})
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", userAgent)

	if forwardedFor, ok := data.Headers["x-forwarded-for"]; ok {
		req.Header.Add("X-Forwarded-For", forwardedFor)
	}

	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(data)
	if err != nil {
		return nil, err
	}

	req.SetBody(body.Bytes())
	return req, nil
}
