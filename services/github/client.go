package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/JKhawaja/rest-example/services"
)

// RealClient ...
type RealClient struct {
	client         *http.Client
	retryPolicy    services.RetryPolicy
	circuitBreaker services.Breaker
	status         bool
}

// NewClient ...
func NewClient() Client {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 10 * time.Second,
	}

	policy := services.NewConstantRetryPolicy(100*time.Millisecond, 3)
	breaker := services.NewBreaker(services.DefaultBreakerConfig)

	return &RealClient{
		client:         &http.Client{Transport: tr},
		retryPolicy:    policy,
		circuitBreaker: breaker,
		status:         true,
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

// GetStatus returns whether or not the backing service is down or not
func (g *RealClient) GetStatus() bool {
	// TODO: if g.status is false, make a health check call on the service, if not 200 then return g.status
	// else change g.status to true, then return it
	return g.status
}

// SetStatus allows the client to specify if the backing service is down or not
// NOTE: changing the status for a http client shared between go-routines will cause a data-race
func (g *RealClient) SetStatus(status bool) {
	g.status = status
}

// ListKeys ...
func (g *RealClient) ListKeys(username string) ([]Key, error) {
	emptyResp := []Key{}

	// TODO: figure out how to use retryPolicies and CircuitBreakers
	// with methods like this one ... need to check the response codes, etc.
	// before trying to json decode the response.
	url := fmt.Sprintf("http://api.github.com/users/%s/keys", username)
	resp, err := g.client.Get(url)
	if err != nil {
		return emptyResp, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Error reading GitHub response body: %+v for username: %s", err, username)
		return emptyResp, err
	}

	var response []Key
	err = json.Unmarshal(b, &response)
	if err != nil {
		err := fmt.Errorf("Could not decode response format: %+v, for username: %s", err, username)
		return emptyResp, err
	}

	return response, nil
}
