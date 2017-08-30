package s3

// Client ...
type Client interface {
	Upload(string, string) (string, error)
}
