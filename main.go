//go:generate goagen bootstrap -d github.com/JKhawaja/replicated/design

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JKhawaja/replicated/app"
	. "github.com/JKhawaja/replicated/controllers"
	"github.com/JKhawaja/replicated/services/github"
	"github.com/JKhawaja/replicated/services/kubernetes"
	"github.com/JKhawaja/replicated/services/s3"
	. "github.com/JKhawaja/replicated/util"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/logrus"
	"github.com/goadesign/goa/middleware"
	"github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
)

func main() {
	flag.Parse()

	if kubernetes.Deploy {
		log.Println("Building binary ...")
		path, err := BuildBinary("rest-example")
		if err != nil {
			log.Fatalf("Could not build Go binary: %s", err.Error())
		}

		// TODO: specify your S3 bucket
		// TODO: add aws keys and s3 bucket is 'kubernetes-test-1234'
		s3Bucket := ""

		client := s3.NewS3Client()
		log.Println("Uploading binary ...")
		url, err := client.Upload(s3Bucket, path)
		if err != nil {
			log.Fatalf("Could not upload binary to S3: %s", err.Error())
		}

		// TODO: specify where your KubeConfig file is
		// $HOME/.kube/config check with `cat config` -- set up for minikube
		kubeConfigPath := "$HOME/.kube/config"

		k8, err := kubernetes.NewKubernetesClient(kubeConfigPath)
		if err != nil {
			log.Fatalf("Could not create a Kubernetes Client: %s", err.Error())
		}

		config := kubernetes.DeploymentConfig{
			BURL: url,
			Name: "rest-example",
		}

		log.Println("Deploying binary ...")
		if err = k8.CreateDeployment(config); err != nil {
			log.Fatalf("Could not create Kubernetes Deployment: %s", err.Error())
		}

		log.Printf("Binary %s successfully deployed to Kubernetes. Please verify with kubectl.", path)
		os.Exit(0)
	}

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
