package commitlog

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNewMinIODispatcher is to ensure MinIO dispatch
func TestNewMinIODispatcher(t *testing.T) {
	opts := dispatcherOptions{"minio"}
	dispatcher, err := newDispatcher(opts)
	require.Nil(t, err)
	log, err := os.OpenFile("/tmp/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	err = dispatcher.Dispatch(log, "id")
	require.Nil(t, err)
}
