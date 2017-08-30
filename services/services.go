package services

import (
	"github.com/JKhawaja/replicated/services/github"
	"github.com/JKhawaja/replicated/services/kubernetes"
)

// Services ...
type Services interface {
	github.Client
	kubernetes.Client
}
