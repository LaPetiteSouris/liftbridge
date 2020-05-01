package remotestorage

import "os"

//LogDispatcher is the remote-storage dispatcher interface for log messages that
// exceed retention policy
type LogDispatcher interface {
	Dispatch(string, string) error
}

// LogRetriever is the remote commitlog retriever interface for log
// that are stored on remote storage
type LogRetriever interface {
	Retrieve(string, int64) (*os.File, *os.File, error)
}
