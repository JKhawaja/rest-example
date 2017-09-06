package test

import (
	"context"
	"net/http"
	"time"

	"github.com/JKhawaja/rest-example/app"
	. "github.com/JKhawaja/rest-example/controllers"

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

	// Mount "keys" controller
	c := NewKeysController(service)
	app.MountKeysController(service, c)

	return &graceful.Server{
		Timeout: time.Duration(15) * time.Second,
		Server:  &http.Server{Addr: ":8080", Handler: service.Mux},
	}
}
