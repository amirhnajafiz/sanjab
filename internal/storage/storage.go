package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Storage manages ceph cluster operations
type Storage struct {
	conn   *s3.S3
	bucket string
}

// NewConnection returns a ceph connection
func NewConnection(cfg Config) (*Storage, error) {
	// create an AWS session with static credentials
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(cfg.Host),
		Credentials:      credentials.NewStaticCredentials(cfg.Access, cfg.Secret, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}

	return &Storage{
		conn:   s3.New(sess),
		bucket: cfg.Bucket,
	}, nil
}

func (s Storage) Upload(name, path string) error {
	// open the local file for reading
	localFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open local file: %v", err)
	}
	defer func() {
		_ = localFile.Close()

		if er := os.Remove(path); er != nil {
			log.Println(fmt.Errorf("failed to remove file %s: %v", path, err))
		}
	}()

	// upload the file to the S3 bucket
	_, err = s.conn.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(name),
		Body:   localFile,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to Ceph: %v", err)
	}

	return nil
}
