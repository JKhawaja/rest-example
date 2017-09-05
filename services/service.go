package services

import "net/http"

// Service is the standard utility interface for all services ...
type Service interface {
	SetRetryPolicy(RetryPolicy)
	SetCircuitBreaker(Breaker)
	SetTransport(*http.Transport)
	GetStatus() bool
	SetStatus(bool)
}
