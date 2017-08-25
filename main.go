//go:generate goagen bootstrap -d github.com/JKhawaja/replicated/design

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/JKhawaja/replicated/app"
	. "github.com/JKhawaja/replicated/controllers"
	"github.com/JKhawaja/replicated/services/github"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/logrus"
	"github.com/goadesign/goa/middleware"
	"github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
)

func main() {
	// Create service
	service := goa.New("GitHub SSH Keys")

	// Initialize logger handler using logrus package
	logger := logrus.New()
	service.WithLogger(goalogrus.New(logger))

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// GitHub Client
	ghc := github.NewGHC()

	// Mount "keys" controller
	c := NewKeysController(service, ghc)
	app.MountKeysController(service, c)

	// Setup graceful server
	server := &graceful.Server{
		Timeout: time.Duration(15) * time.Second,
		Server:  &http.Server{Addr: ":8080", Handler: service.Mux},
	}

	// Start Server
	log.Fatal(server.ListenAndServe())
}
