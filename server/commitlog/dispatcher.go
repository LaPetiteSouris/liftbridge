package commitlog

import (
	"github.com/pkg/errors"

	"github.com/liftbridge-io/liftbridge/server/logger"
	"github.com/liftbridge-io/liftbridge/server/remotestorage"
)

type dispatcherOptions struct {
	Logger  logger.Logger
	storage string
}

// newDeleteCleaner returns a new cleaner which enforces log retention
// policies by deleting segments.
func newDispatcher(opts dispatcherOptions) (remotestorage.LogDispatcher, error) {
	if opts.storage == "minio" {
		return &remotestorage.MinIODispatcherAdapter{}, nil
	}
	return nil, errors.New("remotestorage type is not provided")
}
