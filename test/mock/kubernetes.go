package mock

import (
	"log"

	"github.com/JKhawaja/rest-example/services/kubernetes"
)

// KubernetesClient ...
type KubernetesClient struct{}

// CreateDeployment ...
func (m *KubernetesClient) CreateDeployment(config kubernetes.DeploymentConfig) error {
	log.Println(config)
	return nil
}
