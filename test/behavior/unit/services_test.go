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
	policy2 := services.NewSingleRetryPolicy(100 * time.Millisecond)
	mockService.SetRetryPolicy(policy2)
	breaker2 := services.NewBreaker(services.BreakerConfig{
		ErrorThreshold:   len(policy2.Backoffs()),
		SuccessThreshold: 1,
		Timeout:          10 * time.Second,
	})
	mockService.SetCircuitBreaker(breaker2)

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
	policy3 := services.NewConstantRetryPolicy(100*time.Millisecond, 5)
	mockService.SetRetryPolicy(policy3)
	breaker3 := services.NewBreaker(services.BreakerConfig{
		ErrorThreshold:   len(policy3.Backoffs()),
		SuccessThreshold: 1,
		Timeout:          10 * time.Second,
	})
	mockService.SetCircuitBreaker(breaker3)

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
	policy4 := services.NewExponentialRetryPolicy(services.DefaultBackoffConfig)
	mockService.SetRetryPolicy(policy4)
	breaker4 := services.NewBreaker(services.BreakerConfig{
		ErrorThreshold:   len(policy4.Backoffs()),
		SuccessThreshold: 1,
		Timeout:          10 * time.Second,
	})
	mockService.SetCircuitBreaker(breaker4)

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
	policy5 := services.NewExponentialRetryPolicy(services.DefaultBackoffConfig)
	mockService.SetRetryPolicy(policy5)
	breaker5 := services.NewBreaker(services.BreakerConfig{
		ErrorThreshold:   len(policy5.Backoffs()),
		SuccessThreshold: 1,
		Timeout:          10 * time.Second,
	})
	mockService.SetCircuitBreaker(breaker5)

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
