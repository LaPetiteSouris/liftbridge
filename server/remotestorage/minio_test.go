package remotestorage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestInitMinioClient is to ensure MinIO initialization
func TestInitMinioClient(t *testing.T) {
	minIOAdapter := MinIODispatcherAdapter{}
	err := minIOAdapter.New("127.0.0.1:9000", "minioadmin", "minioadmin", "test", "", false)
	require.Nil(t, err)
}

// TestMakeBucket is to ensure MinIO initialization
func TestMakeBucket(t *testing.T) {
	minIOAdapter := MinIODispatcherAdapter{}
	err := minIOAdapter.New("127.0.0.1:9000", "minioadmin", "minioadmin", "test", "", false)
	require.Nil(t, err)
	err = minIOAdapter.MakeBucket()
	require.Nil(t, err)
}

// TestMakeBucket is to ensure MinIO initialization
func TestDispatch(t *testing.T) {
	minIOAdapter := MinIODispatcherAdapter{}
	err := minIOAdapter.New("127.0.0.1:9000", "minioadmin", "minioadmin", "test", "", false)
	require.Nil(t, err)
	err = minIOAdapter.MakeBucket()
	require.Nil(t, err)
	log, err := os.OpenFile("/tmp/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	err = minIOAdapter.Dispatch(log, "id")
	require.Nil(t, err)
}
