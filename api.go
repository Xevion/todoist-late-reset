package main

import (
	"io"
	"net/http"
	"net/url"
)

func getRecentlyCompleted() (string, error) {
	baseURL := "https://api.todoist.com/sync/v9/activity/get"
	params := url.Values{}
	params.Add("event_type", "completed")

	fullURL := baseURL + "?" + params.Encode()

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+todoistApiToken)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
