package github

import "github.com/JKhawaja/rest-example/services"

// Key is the type for a public SSH key from Github
type Key struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

// Keys ...
type Keys []Key

// Client defines the standard interface for defining GitHub API access functions
type Client interface {
	services.Service
	ListKeys(username string) ([]Key, error)
	HealthCheck() (bool, error)
}
