package client

import "net/http"

type RoundTripper struct {
	token     string
	transport http.RoundTripper
}

func (r *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+r.token)

	return r.transport.RoundTrip(req)
}
