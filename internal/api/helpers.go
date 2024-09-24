package api

import "net/http"

// applyAuthorization sets the Authorization header with the API token.
func (sc *SyncClient) applyAuthorization(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+sc.ApiToken)
}
