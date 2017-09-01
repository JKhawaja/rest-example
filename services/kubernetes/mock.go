package kubernetes

import "log"

// MockClient ...
type MockClient struct{}

// CreateDeployment ...
func (m *MockClient) CreateDeployment(config DeploymentConfig) error {
	log.Println(config)
	return nil
}
