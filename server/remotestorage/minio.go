package remotestorage

import (
	"github.com/minio/minio-go/v6"
)

// MinIODispatcherAdapter is the adapter for dispatching logs to MinIO
type MinIODispatcherAdapter struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
	bucket          string
	location        string
	client          *minio.Client
	LogDispatcher
}

// New generate new MinIO client
func (m *MinIODispatcherAdapter) New(endpoint string, accessKeyID string,
	secretAccessKey string, bucket string, location string, useSSL bool) error {
	m.endpoint = endpoint
	m.accessKeyID = accessKeyID
	m.secretAccessKey = secretAccessKey
	m.useSSL = useSSL
	m.bucket = bucket
	m.location = location

	// Initialize minio client object.
	minioClient, err := minio.New(m.endpoint, m.accessKeyID, m.secretAccessKey, m.useSSL)
	if err != nil {
		return err
	}

	m.client = minioClient
	return nil
}

// MakeBucket try to create the bucket
func (m *MinIODispatcherAdapter) MakeBucket() error {

	err := m.client.MakeBucket(m.bucket, m.location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := m.client.BucketExists(m.bucket)
		if errBucketExists == nil && exists {
			return nil
		}
		return err
	}
	return nil
}

// Dispatch implements LogDispatcher.Dispatch method for MinIO
func (m *MinIODispatcherAdapter) Dispatch(path string, id string) error {
	// Make bucket in case does not exist
	m.MakeBucket()
	_, err := m.client.FPutObject(m.bucket, id, path, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	return nil
}
