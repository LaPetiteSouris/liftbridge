package remotestorage

// MinIODispatcherAdapter is the adapter for dispatching logs to MinIO
type MinIODispatcherAdapter struct {
	d *LogDispatcher
}

// Dispatch implements LogDispatcher.Dispatch method for MinIO
func (*MinIODispatcherAdapter) Dispatch(data interface{}) error {
	return nil
	//
}
