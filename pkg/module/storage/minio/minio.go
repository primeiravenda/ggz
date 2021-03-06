package minio

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"
	"github.com/rs/zerolog/log"
)

// Minio client
type Minio struct {
	host   string
	client *minio.Client
}

// NewEngine struct
func NewEngine(endpoint, accessID, secretKey string, ssl bool) (*Minio, error) {
	var client *minio.Client
	var err error
	if endpoint == "" {
		return nil, errors.New("endpoint can't be empty")
	}

	// Fetching from IAM roles assigned to an EC2 instance.
	if accessID == "" && secretKey == "" {
		iam := credentials.NewIAM("")
		client, err = minio.NewWithCredentials(endpoint, iam, ssl, "")
	} else {
		// Initialize minio client object.
		client, err = minio.New(endpoint, accessID, secretKey, ssl)
	}

	if err != nil {
		return nil, err
	}

	host := ""
	if ssl {
		host = "https://" + endpoint
	} else {
		host = "http://" + endpoint
	}

	return &Minio{
		host:   host,
		client: client,
	}, nil
}

// UploadFile to s3 server
func (m *Minio) UploadFile(bucketName, objectName string, content []byte) error {
	// Upload the zip file with FPutObject
	_, err := m.client.PutObject(
		bucketName,
		objectName,
		bytes.NewReader(content),
		int64(len(content)),
		minio.PutObjectOptions{ContentType: http.DetectContentType(content)},
	)

	return err
}

// CreateBucket create bucket
func (m *Minio) CreateBucket(bucketName, region string) error {
	exists, err := m.client.BucketExists(bucketName)
	if err != nil {
		return err
	}

	if exists {
		log.Info().Msgf("We already own %s bucket", bucketName)
		return nil
	}

	if err := m.client.MakeBucket(bucketName, region); err != nil {
		return err
	}
	log.Info().Msgf("Successfully created s3 bucket: %s", bucketName)

	return nil
}

// FilePath for store path + file name
func (m *Minio) FilePath(_, fileName string) string {
	return fmt.Sprintf("%s/%s", os.TempDir(), fileName)
}

// DeleteFile delete file
func (m *Minio) DeleteFile(bucketName, fileName string) error {
	return m.client.RemoveObject(bucketName, fileName)
}

// GetFile for storage host + bucket + filename
func (m *Minio) GetFile(bucketName, fileName string) string {
	return m.host + "/" + bucketName + "/" + fileName
}
