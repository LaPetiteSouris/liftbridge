package remotestorage

//LogDispatcher is the rem.ote-storage dispatcher interface for log messages that
// exceed retention policy
type LogDispatcher interface {
	Dispatch(string, string) error
}
