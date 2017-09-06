//go:generate goagen bootstrap -d github.com/JKhawaja/rest-example/design

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JKhawaja/rest-example/app"
	. "github.com/JKhawaja/rest-example/controllers"
	"github.com/JKhawaja/rest-example/services/kubernetes"
	"github.com/JKhawaja/rest-example/services/s3"
	. "github.com/JKhawaja/rest-example/util"

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
			log.Fatalf("Could not build Go binary: %+v with value: %s", err, path)
		}

		// AWS S3 bucket and region
		bucket := os.Getenv("AWS_BUCKET")
		region := os.Getenv("AWS_REGION")
		client, err := s3.NewS3Client(region)
		if err != nil {
			log.Fatalf("Could not create an AWS S3 Client: %+v", err)
		}

		log.Println("Uploading binary ...")
		url, err := client.Upload(bucket, path)
		if err != nil {
			log.Fatalf("Could not upload binary to S3: %+v", err)
		}
		log.Printf("Binary uploaded to: %s", url)

		// kubeconfig path
		kubeConfigPath := os.Getenv("KUBE_CONFIG")

		k8, err := kubernetes.NewKubernetesClient(kubeConfigPath)
		if err != nil {
			log.Fatalf("Could not create a Kubernetes Client: %+v", err)
		}

		log.Println("Deploying binary ...")
		if err = k8.CreateDeployment(kubernetes.DeploymentConfig{
			BURL: url,
			Name: "rest-example",
		}); err != nil {
			log.Fatalf("Could not create Kubernetes Deployment: %+v", err)
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

	// Mount "keys" controller
	c := NewKeysController(service)
	app.MountKeysController(service, c)

	// Setup graceful server
	server := &graceful.Server{
		Timeout: time.Duration(15) * time.Second,
		Server:  &http.Server{Addr: ":8080", Handler: service.Mux},
	}

	// Start Server
	log.Fatal(server.ListenAndServe())
}
