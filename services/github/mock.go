package github

import (
	"net/http"

	"github.com/JKhawaja/rest-example/services"

	"github.com/stretchr/testify/mock"
)

// MockClient ...
type MockClient struct {
	mock.Mock
}

// NewMockClient ...
func NewMockClient() Client {
	return &MockClient{}
}

// Keys ...
type Keys []Key

// ListKeys ...
func (g *MockClient) ListKeys(username string) ([]Key, error) {
	args := g.Mock.Called(username)
	return args.Get(0).(Keys), args.Error(1)
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
