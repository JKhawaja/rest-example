package services

import "net/http"

// Service is a standard utility interface that can be embedded in any service interface ...
// NOTE: if an API endpoint requires a different retry-policy or a different circuit-breaker configuration, etc.
// a new client will need to be created inside the handler and the new policies/configs can be specified via
// these methods. This will alter the behavior of the service Client's methods.
type Service interface {
	Do(*http.Request) (*http.Response, error)
	SetRetryPolicy(RetryPolicy)
	SetCircuitBreaker(Breaker)
	SetTransport(*http.Transport)
	GetStatus() bool
	SetStatus(bool)
}
