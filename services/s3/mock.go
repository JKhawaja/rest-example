package s3

import (
	"fmt"
)

// MockS3Client ...
type MockS3Client struct{}

// Upload ...
func (m *MockS3Client) Upload(bucket, path string) error {
	err := fmt.Errorf("%s - %s", bucket, path)
	return err
}
