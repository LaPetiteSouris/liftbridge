package commitlog

import (
	"github.com/pkg/errors"

	"github.com/liftbridge-io/liftbridge/server/remotestorage"
)

type dispatcherOptions struct {
	storage string
}

// newDeleteCleaner returns a new cleaner which enforces log retention
// policies by deleting segments.
func newDispatcher(opts dispatcherOptions) (remotestorage.LogDispatcher, error) {
	if opts.storage == "minio" {
		adapter := &remotestorage.MinIODispatcherAdapter{}
		err := adapter.New("127.0.0.1:9000", "minioadmin", "minioadmin", "test", "", false)
		if err != nil {
			return nil, err
		}
		return adapter, nil

	}
	return nil, errors.New("remotestorage type is not provided")
}
