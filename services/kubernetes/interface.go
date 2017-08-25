package kubernetes

// TODO: Init package with flags to set Configuration

// Configuration ...
type Configuration struct {
}

// Client ...
type Client interface {
	CreateReplicaset(Configuration)
	DeleteReplicaset(Configuration)
}
