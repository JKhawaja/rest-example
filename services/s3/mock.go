package s3

import (
	"fmt"
)

// MockClient ...
type MockClient struct{}

// Upload ...
func (m *MockClient) Upload(bucket, path string) error {
	err := fmt.Errorf("%s - %s", bucket, path)
	return err
}
