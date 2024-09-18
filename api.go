package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	syncToken string
)

type ActivityLog struct {
	Count  int     `json:"count"`
	Events []Event `json:"events"`
}

type Event struct {
	EventDate   time.Time      `json:"event_date"`
	EventType   string         `json:"event_type"`
	ExtraData   map[string]any `json:"extra_data"`
	ExtraDataID int64          `json:"extra_data_id"`
	ID          int64          `json:"id"`

	InitiatorID *int64 `json:"initiator_id"`
	ObjectID    string `json:"object_id"`
	ObjectType  string `json:"object_type"`
}

func getRecentlyCompleted() (*ActivityLog, error) {
	baseURL := "https://api.todoist.com/sync/v9/activity/get"
	params := url.Values{}
	params.Add("event_type", "completed")

	fullURL := baseURL + "?" + params.Encode()

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+todoistApiToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var activityLog *ActivityLog
	if err := json.Unmarshal(body, &activityLog); err != nil {
		return nil, err
	}

	return activityLog, nil
}

// $ curl https://api.todoist.com/sync/v9/sync \
//     -H "Authorization: Bearer 0123456789abcdef0123456789abcdef01234567" \
//     -d sync_token='*' \
//     -d resource_types='["all"]'

func synchronize() {

}
