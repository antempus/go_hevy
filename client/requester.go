package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Requester struct {
	Url        *url.URL
	Headers    map[string]string
	HttpClient *http.Client
	UserAgent  string
	Context    context.Context
}

func (r *Requester) setupRequest(method string, path string, input interface{}, queryParams map[string]string) (*http.Request, error) {
	// TODO: @antempus Determine if this is the correct way to leverage "base path" + "resource path"
	// TODO: Add Logging
	targetUrl := r.Url.RawPath + path
	if queryParams != nil {
		targetUrl = targetUrl + "?"
		for key, value := range queryParams {
			targetUrl = targetUrl + key + "=" + value + "&"
		}
	}

	var buf io.ReadWriter
	if input != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(input)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(r.Context, method, targetUrl, buf)
	if err != nil {
		return nil, err
	}

	if input != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", r.UserAgent)

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}
	return req, nil
}
func (r *Requester) do(req *http.Request, output interface{}) (*http.Response, error) {
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(output)
	return resp, err
}
