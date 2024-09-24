package api

import "encoding/json"

type SyncResponse struct {
	// The new token to use for the next incremental sync.
	SyncToken string `json:"sync_token"`
	// If true, this response is a full sync and the client should not merge but instead replace it's state.
	FullSync bool `json:"full_sync"`
	// Used for commands where local IDs are temporarily chosen by the client and need to be mapped to the server's IDs.
	TempIDMapping map[string]interface{} `json:"temp_id_mapping"`
	Items         []interface{}          `json:"items"`
	// CompletedInfo               []interface{}          `json:"completed_info"`
	// Collaborators               []interface{}          `json:"collaborators"`
	// CollaboratorStates          []interface{}          `json:"collaborator_states"`
	// DayOrders                   map[string]interface{} `json:"day_orders"`
	// Filters                     []interface{}          `json:"filters"`
	// Labels                      []interface{}          `json:"labels"`
	// LiveNotifications           []interface{}          `json:"live_notifications"`
	// LiveNotificationsLastReadID string                 `json:"live_notifications_last_read_id"`
	// Locations                   []interface{}          `json:"locations"`
	// Notes                       []interface{}          `json:"notes"`
	// ProjectNotes                []interface{}          `json:"project_notes"`
	// Projects                    []interface{}          `json:"projects"`
	// Reminders                   []interface{}          `json:"reminders"`
	// Sections                    []interface{}          `json:"sections"`
	// Stats                       map[string]interface{} `json:"stats"`
	// SettingsNotifications       map[string]interface{} `json:"settings_notifications"`
	// User                        map[string]interface{} `json:"user"`
	// UserPlanLimits              map[string]interface{} `json:"user_plan_limits"`
	// UserSettings                map[string]interface{} `json:"user_settings"`
}

type State struct {
	Items map[string]Item
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
func (sc *SyncClient) Synchronize(full bool) (*Changes, error) {
	if sc.RequireFullSync {
		sc.syncToken = "*"
	}

	body := map[string]interface{}{
		"sync_token": sc.syncToken,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return 0, nil, err
	}

	res, err := sc.post("/sync", nil, jsonBody)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var syncResponse SyncResponse
	if err := json.NewDecoder(res.Body).Decode(&syncResponse); err != nil {
		return nil, err
	}

	sc.syncToken = syncResponse.SyncToken

	// changes := &Changes{
	// 	Added:   []string{},
	// 	Updated: []string{},
	// 	Deleted: []string{},
	// }

	// // Process the items in syncResponse.Items to populate changes
	// for _, item := range syncResponse.Items {
	// 	// Assuming item is a map[string]interface{} and has a "status" field
	// 	itemMap := item.(map[string]interface{})
	// 	if status, ok := itemMap["status"].(string); ok {
	// 		switch status {
	// 		case "added":
	// 			changes.Added = append(changes.Added, itemMap["id"].(string))
	// 		case "updated":
	// 			changes.Updated = append(changes.Updated, itemMap["id"].(string))
	// 		case "deleted":
	// 			changes.Deleted = append(changes.Deleted, itemMap["id"].(string))
	// 		}
	// 	}
	// }

	return changes, nil
}
