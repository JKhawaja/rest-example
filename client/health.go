// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "GitHub SSH Keys": health Resource Client
//
// Command:
// $ goagen
// --design=github.com/JKhawaja/rest-example/ssot
// --out=$(GOPATH)\src\github.com\JKhawaja\rest-example
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// HealthcheckHealthPath computes a request path to the healthcheck action of health.
func HealthcheckHealthPath() string {

	return fmt.Sprintf("/health")
}

// Returns a 200 if service is available.
func (c *Client) HealthcheckHealth(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewHealthcheckHealthRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewHealthcheckHealthRequest create the request corresponding to the healthcheck action endpoint of the health resource.
func (c *Client) NewHealthcheckHealthRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
