package github

// Key is the type for a public SSH key from Github
type Key struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

// GitHubClient defines the standard interface for defining GitHub API access functions
type Client interface {
	ListKeys(username string) ([]Key, error)
}
