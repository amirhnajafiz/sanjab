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
	// Specify the S3 bucket and object key
	bucketName := "your-bucket-name"
	objectKey := "example.txt"

	// Specify the path to the local file you want to upload
	localFilePath := "/path/to/local/file.txt"

	// Open the local file for reading
	localFile, err := os.Open(localFilePath)
	if err != nil {
		log.Fatalf("Failed to open local file: %v", err)
	}
	defer localFile.Close()

	// Upload the file to the S3 bucket
	_, err = s.conn.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   localFile,
	})
	if err != nil {
		log.Fatalf("Failed to upload file to S3: %v", err)
	}

	fmt.Printf("File '%s' uploaded to S3 object '%s' in bucket '%s'\n", localFilePath, objectKey, bucketName)

	return nil
}
