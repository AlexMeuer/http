package http

import "net/http"

type CustomHeaderTransport struct {
	Headers map[string]string
	Base    http.RoundTripper
}

func (t *CustomHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.Headers {
		req.Header.Add(k, v)
	}
	base := t.Base
	if base == nil {
		base = http.DefaultTransport
	}
	return base.RoundTrip(req)
}
