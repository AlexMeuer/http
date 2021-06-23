package http

import (
	"net/http"
	"testing"
)

type H map[string]string

type RoundTripper struct {
	Invocations []*http.Request
}

func (rt *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.Invocations = append(rt.Invocations, req)
	return &http.Response{
		Status:     "Test Response",
		StatusCode: 222,
	}, nil
}

func harness(headers H) (*http.Client, *CustomHeaderTransport, *RoundTripper) {
	base := &RoundTripper{}
	cht := &CustomHeaderTransport{
		Headers: headers,
		Base:    base,
	}
	client := &http.Client{
		Transport: cht,
	}
	return client, cht, base
}

func TestHeadersSent(t *testing.T) {
	h := H{"X-Foo": "bar", "Lorem": "Ipsum"}
	c, _, b := harness(h)
	c.Get("http://example.com")
	if len(b.Invocations) != 1 {
		t.Fatalf("Incorrect number of base RoundTripper invocations. Expected: %d, Actual: %d", 1, len(b.Invocations))
	}
	verify := func(key string) {
		if b.Invocations[0].Header.Get(key) != h[key] {
			t.Logf("Bad header. Expected: '%s', Actual: '%s'", h[key], b.Invocations[0].Header.Get(key))
			t.Fail()
		}
	}
	verify("X-Foo")
	verify("Lorem")
}
