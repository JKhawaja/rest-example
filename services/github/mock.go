package github

import (
	"fmt"
	"net/http"

	"github.com/JKhawaja/rest-example/services"
)

// MockClient ...
type MockClient struct{}

// NewMockGHC ...
func NewMockGHC() Client {
	return &MockClient{}
}

// ListKeys ...
func (g *MockClient) ListKeys(username string) ([]Key, error) {
	fmt.Println(username)
	response := []Key{}
	return response, nil
}

// SetRetryPolicy allows you to change the default retry-policy for the client
func (g *MockClient) SetRetryPolicy(policy services.RetryPolicy) {}

// SetCircuitBreaker allows you to change the default circuit-breaker for the client
func (g *MockClient) SetCircuitBreaker(breaker services.Breaker) {}

// SetTransport allows you to change the default http-transport for the client
func (g *MockClient) SetTransport(transport *http.Transport) {}

// GetStatus returns whether or not the backing service is down or not
func (g *MockClient) GetStatus() bool {
	return true
}

// SetStatus allows the client to specify if the backing service is down or not
func (g *MockClient) SetStatus(status bool) {}
