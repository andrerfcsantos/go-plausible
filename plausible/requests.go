package plausible

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type apiError struct {
	Error string `json:"error"`
}

func checkAPIResponseForErrors(resp *fasthttp.Response) ([]byte, error) {

	body := resp.Body()
	status := resp.StatusCode()

	if status < 200 || status > 203 {
		var errorJSON apiError
		err := json.Unmarshal(body, &errorJSON)

		if err != nil {
			return body, fmt.Errorf("non-ok code received (%d) from the API", status)
		}

		return body, fmt.Errorf("api error with code %d: %s", status, errorJSON.Error)
	}

	return body, nil
}

func doRequest(client *fasthttp.Client, req *fasthttp.Request) ([]byte, error) {

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	err := client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	body, err := checkAPIResponseForErrors(resp)
	if err != nil {
		return nil, err
	}

	return body, nil
}
