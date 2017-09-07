package mock

import (
	"fmt"
)

// S3Client ...
type S3Client struct{}

// Upload ...
func (m *S3Client) Upload(bucket, path string) error {
	err := fmt.Errorf("%s - %s", bucket, path)
	return err
}
