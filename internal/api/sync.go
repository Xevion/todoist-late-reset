package api

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type SyncResponse struct {
	// The new token to use for the next incremental sync.
	SyncToken string `json:"sync_token,omitempty"`
	// If true, this response is a full sync and the client should not merge but instead replace its state.
	FullSync bool `json:"full_sync,omitempty"`
	// Used for commands where local IDs are temporarily chosen by the client and need to be mapped to the server's IDs.
	TempIDMapping map[string]interface{} `json:"temp_id_mapping,omitempty"`
	Items         []Item                 `json:"items,omitempty"`
	// CompletedInfo               []interface{}          `json:"completed_info,omitempty"`
	// Collaborators               []interface{}          `json:"collaborators,omitempty"`
	// CollaboratorStates          []interface{}          `json:"collaborator_states,omitempty"`
	// DayOrders                   map[string]interface{} `json:"day_orders,omitempty"`
	// Filters                     []interface{}          `json:"filters,omitempty"`
	// Labels                      []interface{}          `json:"labels,omitempty"`
	// LiveNotifications           []interface{}          `json:"live_notifications,omitempty"`
	// LiveNotificationsLastReadID string                 `json:"live_notifications_last_read_id,omitempty"`
	// Locations                   []interface{}          `json:"locations,omitempty"`
	// Notes                       []interface{}          `json:"notes,omitempty"`
	// ProjectNotes                []interface{}          `json:"project_notes,omitempty"`
	// Projects                    []interface{}          `json:"projects,omitempty"`
	// Reminders                   []interface{}          `json:"reminders,omitempty"`
	// Sections                    []interface{}          `json:"sections,omitempty"`
	// Stats                       map[string]interface{} `json:"stats,omitempty"`
	// SettingsNotifications       map[string]interface{} `json:"settings_notifications,omitempty"`
	// User                        map[string]interface{} `json:"user,omitempty"`
	// UserPlanLimits              map[string]interface{} `json:"user_plan_limits,omitempty"`
	// UserSettings                map[string]interface{} `json:"user_settings,omitempty"`
}

type State struct {
	Items map[string]Item
}

// NewState creates a new blank state with initialized fields.
func NewState() *State {
	return &State{
		Items: make(map[string]Item),
	}
}

// Synchronize the client's state with the server. If the full parameter is set to true,
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
func (sc *SyncClient) Synchronize(full bool) (*[]byte, error) {
	if sc.requireFullSync || sc.lastFullSync.IsZero() {
		fmt.Printf("Performing full sync\n")
		sc.syncToken = "*"
	}

	resourceTypes := make([]string, 4)
	for resourceType := range sc.resourceTypes {
		resourceTypes = append(resourceTypes, string(resourceType))
	}

	body := map[string]interface{}{
		"sync_token":     sc.syncToken,
		"resource_types": resourceTypes,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	res, err := sc.post("/sync", nil, jsonBody)
	if err != nil {
		return nil, fmt.Errorf("failed to post sync request: %w", err)
	}
	now := time.Now()

	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var syncResponse SyncResponse
	if err := json.Unmarshal(bodyBytes, &syncResponse); err != nil {
		return nil, fmt.Errorf("failed to decode sync response: %w", err)
	}

	// Always set in every request
	sc.syncToken = syncResponse.SyncToken
	sc.lastSync = now

	// Simple, replace the state with the new items
	if syncResponse.FullSync {
		sc.lastFullSync = now
		sc.requireFullSync = false

		if syncResponse.Items != nil {
			sc.State.Items = make(map[string]Item)
			for _, item := range syncResponse.Items {
				sc.State.Items[item.ID] = item
			}
		}
	} else {
		// Partial sync
		if syncResponse.Items != nil {
			for _, item := range syncResponse.Items {
				sc.State.Items[item.ID] = item

				if item.IsDeleted {
					delete(sc.State.Items, item.ID)
				}
			}
		}
	}

	return nil, nil
	// return changes, nil
}
