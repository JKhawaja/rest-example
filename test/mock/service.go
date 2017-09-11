package mock

import (
	"net/http"
	"time"

	"github.com/JKhawaja/rest-example/services"

	"gopkg.in/eapache/go-resiliency.v1/breaker"
)

// MockService ...
type MockService struct {
	RetryPolicy    services.RetryPolicy
	CircuitBreaker services.Breaker
	Client         *http.Client
	Status         services.Status
}

// NewMockService ...
func NewMockService() services.Service {
	policy := services.NewNoRetryPolicy()
	breaker := services.NewBreaker(services.DefaultBreakerConfig)
	client := &http.Client{}
	status := services.NewStatus()
	status.Set("Service", true)

	return &MockService{
		RetryPolicy:    policy,
		CircuitBreaker: breaker,
		Client:         client,
		Status:         status,
	}
}

// Do ...
func (m *MockService) Do(req *http.Request) (*http.Response, error) {
	response := &http.Response{}
	backoffs := m.RetryPolicy.Backoffs()
	retries := 0

	for {
		result := m.CircuitBreaker.CB.Run(func() error {
			resp, err := m.Client.Do(req)
			response = resp
			if err != nil {
				return err
			}
			return nil
		})

		switch result {
		case nil:
			// Success
			m.SetStatus(true)
			return response, nil
		case breaker.ErrBreakerOpen:
			// Circuit-Breaker open
			m.SetStatus(false)
			return response, breaker.ErrBreakerOpen
		default:
			// Otherwise, retry
			if retries <= len(backoffs) && m.RetryPolicy.Retry(req) {
				time.Sleep(backoffs[retries])
				retries++
			} else {
				return response, breaker.ErrBreakerOpen
			}
		}
	}
}

// SetRetryPolicy ...
func (m *MockService) SetRetryPolicy(policy services.RetryPolicy) {
	m.RetryPolicy = policy
}

// SetCircuitBreaker ...
func (m *MockService) SetCircuitBreaker(breaker services.Breaker) {
	m.CircuitBreaker = breaker
}

// SetTransport ...
func (m *MockService) SetTransport(transport *http.Transport) {
	m.Client.Transport = transport
}

// GetStatus ...
func (m *MockService) GetStatus() bool {
	return m.Status.Get("service")
}

// SetStatus ...
func (m *MockService) SetStatus(status bool) {
	m.Status.Set("service", status)
}
