package services

import (
	"github.com/JKhawaja/rest-example/services/github"
	"github.com/JKhawaja/rest-example/services/kubernetes"
)

// Services ...
type Services interface {
	github.Client
	kubernetes.Client
}
