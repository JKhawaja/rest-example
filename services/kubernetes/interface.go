package kubernetes

// TODO: Init package with flags to set Configuration

// Configuration ...
type Configuration struct {
}

// Client ...
type Client interface {
	CreateDeployment(DeploymentConfig) error
	// DeleteDeployment() error
}
