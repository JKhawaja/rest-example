//go:generate goagen bootstrap -d github.com/JKhawaja/replicated/design

package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JKhawaja/replicated/app"
	. "github.com/JKhawaja/replicated/controllers"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

// GitHubClient is a type-alias for the standard http client ...
type GitHubClient http.Client

func main() {
	// Create service
	service := goa.New("GitHub SSH Keys")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// GitHub Client
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 10 * time.Second,
	}
	ghc := &http.Client{Transport: tr}

	// Mount "keys" controller
	c := NewKeysController(service, ghc)
	app.MountKeysController(service, c)

	// Create Shutdown Channels
	errChan := make(chan error, 10)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start Server
	go func() {
		errChan <- service.ListenAndServe(":8080")
	}()

	// Blocking Clean Shutdown
	for {
		select {
		case err := <-errChan:
			if err != nil {
				service.LogError("startup", "err", err)
				os.Exit(1)
			}
		case s := <-signalChan:
			service.LogError("Crash: %s", s.String())
			os.Exit(0)
		}
	}
}
