package storage

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Storage struct {
	conn *s3.S3
}

// NewConnection returns a ceph connection
func NewConnection(confFile, pool string) (*Storage, error) {
	// Set your Ceph Object Gateway (S3) endpoint, access key, and secret key
	endpoint := "http://your-ceph-rgw-endpoint:port"
	accessKey := "your-access-key"
	secretKey := "your-secret-key"

	// Create an AWS session with static credentials
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(endpoint),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	return &Storage{
		conn: s3.New(sess),
	}, nil
}

func (s Storage) Upload() error {

}
