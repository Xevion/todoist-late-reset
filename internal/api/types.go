package api

import (
	"net/http"
	"time"
)

type SyncClient struct {
	Http         *http.Client
	SyncToken    string
	ApiToken     string
	LastSync     time.Time
	LastFullSync time.Time
}

func NewSyncClient(apiToken string) *SyncClient {
	return &SyncClient{
		Http:      &http.Client{},
		ApiToken:  apiToken,
		SyncToken: "*",
	}
}

type SyncResponse struct {
}

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
