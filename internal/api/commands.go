package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func (sc *SyncClient) RecentlyCompleted() (*ActivityLog, error) {
	baseURL := API_BASE_URL + "/activity/get"
	params := url.Values{}
	params.Add("event_type", "completed")

	fullURL := baseURL + "?" + params.Encode()

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+sc.ApiToken)

	resp, err := sc.Http.Do(req)
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

func (sc *SyncClient) sync() {
	// Implementation for synchronize function
}
