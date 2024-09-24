package api

type ClientState struct {
	// Items map[string]Item
}

type Changes struct {
	Added   []string
	Updated []string
	Deleted []string
}

// sync synchronizes the client's state with the server. If the full parameter is set to true,
// a full synchronization is performed, otherwise, a partial synchronization is done.
// This strongly mutates the client's state.
//
// Parameters:
//
//	full - a boolean indicating whether to perform a full synchronization.
//
// Returns:
//
//	int - the number of changes synchronized.
//	*Changes - a pointer to a Changes struct containing the details of the changes.
//	error - an error object if an error occurred during synchronization, otherwise nil.
func (sc *SyncClient) sync(full bool) (int, *Changes, error) {
	return 0, nil, nil
}
