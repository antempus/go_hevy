package client

import (
	"bytes"
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
}

func (r *Requester) setupRequest(method string, path string, body interface{}) (*http.Request, error) {
	// TODO: @antempus Determine if this is the correct way to leverage "base path" + "resource path"
	// TODO: Add Logging
	targetUrl := r.Url.RawPath + path
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, targetUrl, buf)

	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", r.UserAgent)
	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}
	return req, nil
}
func (r *Requester) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
