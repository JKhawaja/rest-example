package services

import (
	"time"

	"gopkg.in/eapache/go-resiliency.v1/breaker"
)

// BreakerConfig ...
type BreakerConfig struct {
	// NOTE: should set error threshold to the length of the backoffs slice in the RetryPoliicy (!!)
	ErrorThreshold int
	// NOTE: SuccessThreshold should almost always be 1
	SuccessThreshold int
	Timeout          time.Duration
}

// DefaultBreakerConfig ...
var DefaultBreakerConfig = BreakerConfig{
	ErrorThreshold:   3,
	SuccessThreshold: 1,
	Timeout:          10 * time.Second,
}

// Breaker ...
type Breaker struct {
	CB *breaker.Breaker
}

// NewBreaker ...
func NewBreaker(config BreakerConfig) Breaker {
	if config.ErrorThreshold == 0 {
		config.ErrorThreshold = DefaultBreakerConfig.ErrorThreshold
	}
	if config.SuccessThreshold == 0 {
		config.SuccessThreshold = DefaultBreakerConfig.SuccessThreshold
	}
	if config.Timeout == 0 {
		config.Timeout = DefaultBreakerConfig.Timeout
	}

	return Breaker{
		CB: breaker.New(config.ErrorThreshold, config.SuccessThreshold, config.Timeout),
	}
}

/*

	// EXAMPLE:

	for {
		result := b.Run(func() error {
			// communicate with some external service and
			// return an error if the communication failed
			return nil
		})

		switch result {
		case nil:
			// success!
			// set service status to true
		case breaker.ErrBreakerOpen:
			// our function wasn't run because the breaker was open
			// set service status to false
		default:
			// retry based on retryPolicy
		}
	}
*/
