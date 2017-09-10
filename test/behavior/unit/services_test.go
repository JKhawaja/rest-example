// +build unit

package unit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/JKhawaja/rest-example/services"
	"github.com/JKhawaja/rest-example/test/mock"
)

// Intra-Compositional Test for Services package
func TestServices(t *testing.T) {

	mockService := mock.NewMockService()

	// No-Retry Policy
	policy := services.NewNoRetryPolicy()
	mockService.SetRetryPolicy(policy)
	breaker := services.NewBreaker(services.BreakerConfig{
		ErrorThreshold:   len(policy.Backoffs()),
		SuccessThreshold: 1,
		Timeout:          10 * time.Second,
	})
	mockService.SetCircuitBreaker(breaker)

	req, err := http.NewRequest("GET", "localhost:8080", nil)
	if err != nil {
		t.Fatalf("Services Package Test - NoRetryPolicy - Something went wrong with creating a fake request: %+v", err)
	}

	_, err = mockService.Do(req)
	if err.Error() != "circuit breaker is open" {
		t.Fatalf("Services Package Test - NoRetryPolicy - Something went wrong sending the fake request to the non-existent server: %+v", err)
	}

	if mockService.GetStatus() {
		t.Fatal("Services Package Test - NoRetryPolicy - Mock Service should have been marked as down.")
	}

	// Single-Retry Policy
	policy = services.NewSingleRetryPolicy(100 * time.Millisecond)
	mockService.SetRetryPolicy(policy)
	breaker = services.NewBreaker(services.BreakerConfig{
		ErrorThreshold:   len(policy.Backoffs()),
		SuccessThreshold: 1,
		Timeout:          10 * time.Second,
	})
	mockService.SetCircuitBreaker(breaker)

	req, err = http.NewRequest("GET", "localhost:8080", nil)
	if err != nil {
		t.Fatalf("Services Package Test - SingleRetryPolicy - Something went wrong with creating a fake request: %+v", err)
	}

	_, err = mockService.Do(req)
	if err.Error() != "circuit breaker is open" {
		t.Fatalf("Services Package Test - SingleRetryPolicy - Something went wrong sending the fake request to the non-existent server: %+v", err)
	}

	if mockService.GetStatus() {
		t.Fatal("Services Package Test - SingleRetryPolicy - Mock Service should have been marked as down.")
	}

	// Constant-Retry Policy
	policy = services.NewConstantRetryPolicy(100*time.Millisecond, 5)
	mockService.SetRetryPolicy(policy)
	breaker = services.NewBreaker(services.BreakerConfig{
		ErrorThreshold:   len(policy.Backoffs()),
		SuccessThreshold: 1,
		Timeout:          10 * time.Second,
	})
	mockService.SetCircuitBreaker(breaker)

	req, err = http.NewRequest("GET", "localhost:8080", nil)
	if err != nil {
		t.Fatalf("Services Package Test - ConstantRetryPolicy - Something went wrong with creating a fake request: %+v", err)
	}

	_, err = mockService.Do(req)
	if err.Error() != "circuit breaker is open" {
		t.Fatalf("Services Package Test - ConstantRetryPolicy - Something went wrong sending the fake request to the non-existent server: %+v", err)
	}

	if mockService.GetStatus() {
		t.Fatal("Services Package Test - ConstantRetryPolicy - Mock Service should have been marked as down.")
	}

	// Exponential-Retry Policy
	policy = services.NewExponentialRetryPolicy(services.DefaultBackoffConfig)
	mockService.SetRetryPolicy(policy)
	breaker = services.NewBreaker(services.BreakerConfig{
		ErrorThreshold:   len(policy.Backoffs()),
		SuccessThreshold: 1,
		Timeout:          10 * time.Second,
	})
	mockService.SetCircuitBreaker(breaker)

	req, err = http.NewRequest("GET", "localhost:8080", nil)
	if err != nil {
		t.Fatalf("Services Package Test - ExponentialRetryPolicy - Something went wrong with creating a fake request: %+v", err)
	}

	_, err = mockService.Do(req)
	if err.Error() != "circuit breaker is open" {
		t.Fatalf("Services Package Test - ExponentialRetryPolicy - Something went wrong sending the fake request to the non-existent server: %+v", err)
	}

	if mockService.GetStatus() {
		t.Fatal("Services Package Test - ExponentialRetryPolicy - Mock Service should have been marked as down.")
	}

	// Fake Working Service ...
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	req, err = http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatalf("Services Package Test - FakeService - Something went wrong with creating a fake request: %+v", err)
	}

	resp, err := mockService.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("Services Package Test - FakeService - Something went wrong sending the fake request to the fake server: %+v", err)
	}

	if !mockService.GetStatus() {
		t.Fatal("Services Package Test - FakeService - Mock Service should have been marked as up.")
	}
}
