package remotestorage

import (
	"io"

	"github.com/minio/minio-go/v6"
)

// MinIODispatcherAdapter is the adapter for dispatching logs to MinIO
type MinIODispatcherAdapter struct {
	client *minio.Client
	LogDispatcher
}

// New generate new MinIO client
func (m *MinIODispatcherAdapter) New() error {
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		return err
	}

	m.client = minioClient
	return nil
}

// MakeBucket try to create the bucket
func (m *MinIODispatcherAdapter) MakeBucket() error {
	bucketName := "mymusic"
	location := "us-east-1"

	err := m.client.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := m.client.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			return nil
		}
		return err
	}
	return nil
}

// Dispatch implements LogDispatcher.Dispatch method for MinIO
func (m *MinIODispatcherAdapter) Dispatch(data io.Reader) error {
	dataStat, err := data.Stat()
	if err != nil {
		return err
	}
	err = m.client.PutObject("mybucket", "myobect", data, dataStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	return nil
	//
}
