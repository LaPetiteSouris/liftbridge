package remotestorage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestInitMinioClient is to ensure MinIO initialization
func TestInitMinioClient(t *testing.T) {
	minIOAdapter := MinIODispatcherAdapter{}
	err := minIOAdapter.New()
	require.Nil(t, err)
}

// TestMakeBucket is to ensure MinIO initialization
func TestMakeBucket(t *testing.T) {
	minIOAdapter := MinIODispatcherAdapter{}
	err := minIOAdapter.New()
	require.Nil(t, err)
	err = minIOAdapter.MakeBucket()
	require.Nil(t, err)
}

// TestMakeBucket is to ensure MinIO initialization
func TestDispatch(t *testing.T) {
	minIOAdapter := MinIODispatcherAdapter{}
	err := minIOAdapter.New()
	require.Nil(t, err)
	err = minIOAdapter.MakeBucket()
	require.Nil(t, err)

	// New Segment

	err = minIOAdapter.Dispatch()
}
