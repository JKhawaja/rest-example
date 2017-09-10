package mock

import (
	"net/http"

	"github.com/JKhawaja/rest-example/services"
	"github.com/JKhawaja/rest-example/services/github"

	"github.com/stretchr/testify/mock"
)

// GithubClient ...
type GithubClient struct {
	mock.Mock
}

// NewGithubClient ...
func NewGithubClient() github.Client {
	return &GithubClient{}
}

// ListKeys ...
func (g *GithubClient) ListKeys(username string) ([]github.Key, error) {
	args := g.Mock.Called(username)
	return args.Get(0).([]github.Key), args.Error(1)
}

// HealthCheck ...
func (g *GithubClient) HealthCheck() (bool, error) {
	return true, nil
}

// Do ...
func (g *GithubClient) Do(req *http.Request) (*http.Response, error) {
	resp := &http.Response{
		Request: req,
	}

	return resp, nil
}

// SetRetryPolicy allows you to change the default retry-policy for the client
func (g *GithubClient) SetRetryPolicy(policy services.RetryPolicy) {}

// SetCircuitBreaker allows you to change the default circuit-breaker for the client
func (g *GithubClient) SetCircuitBreaker(breaker services.Breaker) {}

// SetTransport allows you to change the default http-transport for the client
func (g *GithubClient) SetTransport(transport *http.Transport) {}

// GetStatus returns whether or not the backing service is down or not
func (g *GithubClient) GetStatus() bool {
	return true
}

// SetStatus allows the client to specify if the backing service is down or not
func (g *GithubClient) SetStatus(status bool) {}
