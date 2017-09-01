package github

import (
	"fmt"
)

// MockClient ...
type MockClient struct{}

// NewMockGHC ...
func NewMockGHC() Client {
	return &MockClient{}
}

// ListKeys ...
func (g *MockClient) ListKeys(username string) ([]Key, error) {
	fmt.Println(username)
	response := []Key{}
	return response, nil
}
