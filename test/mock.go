package test

import (
	"context"
	"net/http"
	"time"

	"github.com/JKhawaja/rest-example/app"
	. "github.com/JKhawaja/rest-example/controllers"
	"github.com/JKhawaja/rest-example/test/mock"

	"github.com/goadesign/goa"
	"github.com/tylerb/graceful"
)

// Case ...
type Case struct {
	Context context.Context
	Path    string
	Payload []string
}

// NewMockServer ...
func NewMockServer() *graceful.Server {
	// Service
	service := goa.New("Mock Service")

	// Mock GitHub Client
	mockGH := mock.NewGithubClient()

	// Mount "keys" controller
	c := NewKeysController(service, mockGH)
	app.MountKeysController(service, c)

	return &graceful.Server{
		Timeout: time.Duration(15) * time.Second,
		Server:  &http.Server{Addr: ":9090", Handler: service.Mux},
	}
}
