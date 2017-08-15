package services

import (
	"github.com/JKhawaja/replicated/services/github"
)

// Services ...
type Services interface {
	github.Client
}
