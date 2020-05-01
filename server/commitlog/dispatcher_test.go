package commitlog

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

// TestNewMinIODispatcher is to ensure MinIO dispatch
func TestNewMinIODispatcher(t *testing.T) {
	opts := dispatcherOptions{"minio"}
	dispatcher, err := newDispatcher(opts)
	require.Nil(t, err)

	tmpFile := createTempFile()
	fileID := strings.Trim(tmpFile, "/")
	defer os.Remove(tmpFile)

	err = dispatcher.Dispatch(tmpFile, fileID)
	require.Nil(t, err)
}
