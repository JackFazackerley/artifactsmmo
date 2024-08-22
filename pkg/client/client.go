package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://api.artifactsmmo.com/"

type Client struct {
	client *http.Client

	url *url.URL
}

func NewClient(token string) (Client, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return Client{}, fmt.Errorf("parsing URL: %w", err)
	}

	return Client{
		client: &http.Client{
			Transport: &RoundTripper{
				token:     token,
				transport: http.DefaultTransport,
			},
			Timeout: time.Second,
		},
		url: url,
	}, nil
}

func (c Client) Do(ctx context.Context, path, method string, payload any, v any) error {
	url, err := c.url.Parse(path)
	if err != nil {
		return fmt.Errorf("parsing url: %w", err)
	}

	body := bytes.NewBuffer(nil)

	if payload != nil {
		if err := json.NewEncoder(body).Encode(payload); err != nil {
			return fmt.Errorf("encoding payload: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url.String(), body)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse

		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return fmt.Errorf("decoding body: %w", err)
		}

		return errorResponse.Error
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("decoding body: %w", err)
	}

	return nil
}
