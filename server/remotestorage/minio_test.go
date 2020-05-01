package remotestorage

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTempFile() string {
	file, err := ioutil.TempFile(os.TempDir(), "miniotest")
	if err != nil {
		log.Fatal(err)
	}
	return file.Name()
}

// TestInitMinioClient is to ensure MinIO initialization
func TestInitMinioClient(t *testing.T) {
	minIOAdapter := MinIOAdapter{}
	err := minIOAdapter.New("127.0.0.1:9000", "minioadmin", "minioadmin", "test", "", false)
	require.Nil(t, err)
}

// TestMakeBucket is to ensure MinIO initialization
func TestMakeBucket(t *testing.T) {
	minIOAdapter := MinIOAdapter{}
	err := minIOAdapter.New("127.0.0.1:9000", "minioadmin", "minioadmin", "test", "", false)
	require.Nil(t, err)
	err = minIOAdapter.MakeBucket()
	require.Nil(t, err)
}

// TestMakeBucket is to ensure MinIO initialization
func TestDispatch(t *testing.T) {
	minIOAdapter := MinIOAdapter{}
	err := minIOAdapter.New("127.0.0.1:9000", "minioadmin", "minioadmin", "test", "", false)
	require.Nil(t, err)
	err = minIOAdapter.MakeBucket()
	require.Nil(t, err)
	tmpFile := createTempFile()
	fileID := strings.Trim(tmpFile, "/")
	defer os.Remove(tmpFile)
	err = minIOAdapter.Dispatch(tmpFile, fileID)
	require.Nil(t, err)
}
