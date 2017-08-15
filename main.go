//go:generate goagen bootstrap -d github.com/JKhawaja/replicated/design

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/JKhawaja/replicated/app"
	. "github.com/JKhawaja/replicated/controllers"
	"github.com/JKhawaja/replicated/services/github"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/logrus"
	"github.com/goadesign/goa/middleware"
	"github.com/sirupsen/logrus"
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

	// Create Shutdown Channels
	errChan := make(chan error, 10)
	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start Server
	go func() {
		errChan <- service.ListenAndServe(":8080")
	}()

	// Blocking Clean Shutdown (will not work on Windows)
	for {
		select {
		case err := <-errChan:
			if err != nil {
				service.LogError("startup", "err", err)
				os.Exit(1)
			}
		case s := <-signalChan:
			service.LogInfo("Terminate: %s", s.String())
			os.Exit(0)
		}
	}

}
