package remotestorage

import (
	"io"
)

//LogDispatcher is the rem.ote-storage dispatcher interface for log messages that
// exceed retention policy
type LogDispatcher interface {
	Dispatch(io.Reader) error
}
