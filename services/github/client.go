package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/JKhawaja/rest-example/services"

	"gopkg.in/eapache/go-resiliency.v1/breaker"
)

// RealClient ...
type RealClient struct {
	client         *http.Client
	retryPolicy    services.RetryPolicy
	circuitBreaker services.Breaker
	status         services.Status
}

// Health ...
type Health struct {
	Status  string `json:"status"`
	Updated string `json:"last_updated"`
}

// Error ...
type Error struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}

// GHError ...
type GHError struct {
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
	Error   error
}

// NewClient ...
func NewClient(status services.Status) Client {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 10 * time.Second,
	}

	policy := services.NewConstantRetryPolicy(100*time.Millisecond, 3)
	breaker := services.NewBreaker(services.DefaultBreakerConfig)

	realClient := &RealClient{
		client:         &http.Client{Transport: tr},
		retryPolicy:    policy,
		circuitBreaker: breaker,
		status:         status,
	}

	realClient.SetStatus(true)

	return realClient
}

// Do ...
func (g *RealClient) Do(req *http.Request) (*http.Response, error) {
	response := &http.Response{}
	b := g.circuitBreaker.CB
	backoffs := g.retryPolicy.Backoffs()
	retries := 0
	for {
		result := b.Run(func() error {
			resp, err := g.client.Do(req)
			response = resp
			if err != nil {
				return err
			}
			return nil
		})

		switch result {
		case nil:
			// Success
			g.SetStatus(true)
			return response, nil
		case breaker.ErrBreakerOpen:
			// Circuit-Breaker open
			g.SetStatus(false)
			return response, breaker.ErrBreakerOpen
		default:
			// Otherwise, retry
			if retries <= len(backoffs) && g.retryPolicy.Retry(req) {
				time.Sleep(backoffs[retries])
				retries++
			} else {
				return response, breaker.ErrBreakerOpen
			}
		}
	}
}

// SetRetryPolicy allows you to change the default retry-policy for the client
// NOTE: changing the retry-policy for a http client shared between go-routines will cause a data-race
func (g *RealClient) SetRetryPolicy(policy services.RetryPolicy) {
	g.retryPolicy = policy
}

// SetCircuitBreaker allows you to change the default circuit-breaker for the client
// NOTE: changing the circuit-breaker for a http client shared between go-routines will cause a data-race
func (g *RealClient) SetCircuitBreaker(breaker services.Breaker) {
	g.circuitBreaker = breaker
}

// SetTransport allows you to change the default http-transport for the client
// NOTE: changing the http-transport for a http client shared between go-routines will cause a data-race
func (g *RealClient) SetTransport(transport *http.Transport) {
	g.client = &http.Client{Transport: transport}
}

// GetStatus ...
func (g *RealClient) GetStatus() bool {
	if !g.status.Get("github") {
		healthy, err := g.HealthCheck()
		if err.Error != nil {
			return false
		}

		if healthy {
			g.SetStatus(true)
		} else {
			return false
		}
	}

	return true
}

// SetStatus ...
func (g *RealClient) SetStatus(status bool) {
	g.status.Set("github", status)
}

// ListKeys ...
func (g *RealClient) ListKeys(username string) ([]Key, GHError) {
	var response []Key
	errorResponse := GHError{}

	url := fmt.Sprintf("http://api.github.com/users/%s/keys", username)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		errorResponse.Error = err
		return response, errorResponse
	}

	resp, err := g.Do(req)
	if err != nil {
		errorResponse.Error = err
		return response, errorResponse
	}

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			errorResponse.Error = err
			return response, errorResponse
		}
	} else {
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			errorResponse.Error = err
			return response, errorResponse
		}
		return response, errorResponse
	}

	return response, errorResponse
}

// HealthCheck ...
func (g *RealClient) HealthCheck() (bool, GHError) {
	// TODO: retry-policy and circuit-breaker pattern

	errorResponse := GHError{}

	url := "https://status.github.com/api/status.json"
	resp, err := g.client.Get(url)
	if err != nil {
		errorResponse.Error = err
		return g.GetStatus(), errorResponse
	}
	defer resp.Body.Close()

	var health Health
	err = json.NewDecoder(resp.Body).Decode(&health)
	if err != nil {
		errorResponse.Error = err
		return g.GetStatus(), errorResponse
	}

	if health.Status != "good" || health.Status != "minor" {
		g.SetStatus(false)
		return false, errorResponse
	}

	g.SetStatus(true)

	return true, errorResponse
}
