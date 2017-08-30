package s3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3Client ...
type S3Client struct {
	Uploader *s3manager.Uploader
}

// NewS3Client ...
func NewS3Client() Client {
	// Create a Session with a custom region
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	client := s3.New(sess)

	uploader := s3manager.NewUploaderWithClient(client, func(u *s3manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024 // 10MB per part
	})

	return &S3Client{
		Uploader: uploader,
	}
}

// Upload ...
func (s *S3Client) Upload(bucket string, path string) (string, error) {
	// Open file
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	key := f.Name()

	// upload parameters
	params := &s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   f,
	}

	// Upload file
	result, err := s.Uploader.Upload(params)
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
