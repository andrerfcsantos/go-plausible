package plausible

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/valyala/fasthttp"
)

type EventRecorder struct {
	domainName string
	client     *fasthttp.Client
	baseURL    string
}

type eventRequest struct {
	// the domain of your site (must match exactly)
	Domain string `json:"domain"`

	// the name of the event, e.g. 'pageview', otherwise a custom event
	Name string `json:"name"`

	// the URL of the current request
	URL *url.URL `json:"url"`
}

func NewEventRecorder(domainName string, client *fasthttp.Client, baseURL ...string) *EventRecorder {
	var u string
	if len(baseURL) == 0 {
		u = DefaultBaseURL
	} else {
		u = baseURL[0]
	}

	return &EventRecorder{
		domainName: domainName,
		client:     client,
		baseURL:    u,
	}
}

func (e *EventRecorder) Push(r *http.Request, event string) error {
	data := eventRequest{
		Domain: e.domainName,
		Name:   event,
		URL:    r.URL,
	}

	var userAgent string
	if r.Header.Get("User-Agent") == "" {
		return fmt.Errorf("missing User-Agent header in request")
	}
	userAgent = r.Header.Get("User-Agent")

	forwardedFor := ""
	if r.Header.Get("X-Forwarded-For") != "" {
		forwardedFor = r.Header.Get("X-Forwarded-For")
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(e.baseURL + "/api/event")
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.Add("Content-Type", "application/json")

	// forward request headers to Plausible
	req.Header.Set("X-Forwarded-For", forwardedFor)
	req.Header.Set("User-Agent", userAgent)

	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(data)
	if err != nil {
		return err
	}

	req.SetBody(body.Bytes())
	_, err = doRequest(e.client, req)
	return err
}
