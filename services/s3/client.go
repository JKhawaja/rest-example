package s3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// RealClient ...
type RealClient struct {
	Uploader *s3manager.Uploader
}

// NewS3Client ...
func NewS3Client(region string) (Client, error) {
	var emptyClient RealClient
	creds := credentials.NewEnvCredentials()

	_, err := creds.Get()
	if err != nil {
		return &emptyClient, err
	}

	endpoint := "s3.amazonaws.com"
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      &region,
		Endpoint:    &endpoint,
		Credentials: creds,
	}))
	client := s3.New(sess)

	uploader := s3manager.NewUploaderWithClient(client, func(u *s3manager.Uploader) {
		u.PartSize = 64 * 1024 * 1024 // 64MB per part
	})

	return &RealClient{
		Uploader: uploader,
	}, nil
}

// Upload ...
func (s *RealClient) Upload(bucket string, path string) (string, error) {
	// Open file
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	key := f.Name()

	// allow binary to be downloadable publicly
	acl := "public-read"
	// upload parameters
	params := &s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   f,
		ACL:    &acl,
	}

	// Upload file
	result, err := s.Uploader.Upload(params)
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
