package api

import (
	"encoding/json"
	"io"
	"net/url"
	"time"
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

func (sc *SyncClient) RecentlyCompleted() (*ActivityLog, error) {
	params := url.Values{"event_type": {"completed"}}

	resp, err := sc.get("activity/get", params)
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
