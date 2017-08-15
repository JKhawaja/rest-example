package github

import (
	"fmt"
)

// MockGHC ...
type MockGHC struct{}

// NewMockGHC ...
func NewMockGHC() Client {
	return &MockGHC{}
}

// ListKeys ...
func (g *MockGHC) ListKeys(username string) ([]Key, error) {
	fmt.Println(username)
	response := []Key{}
	return response, nil
}
