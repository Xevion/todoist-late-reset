package api

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

var (
	userAgent string
)

func init() {
	revision := "unknown"
	version := "v0.0.0"
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "(devel)" {
			version = info.Main.Version
			revision = info.Main.Sum
		} else {
			fmt.Println("WARN : Inaccurate version information")
		}
	}
	userAgent = fmt.Sprintf("todoist-late-reset/%v (%v; revision %v)", version, runtime.GOOS, revision)
	fmt.Println(userAgent)
}

// SyncClient represents a client for synchronizing data with the Todoist API.
// It holds the HTTP client, synchronization tokens, timestamps of the last syncs,
// and the types of resources to be synchronized.
type SyncClient struct {
	// lastSync is the timestamp of the last synchronization, full or incremental.
	lastSync time.Time
	// lastFullSync is the timestamp of the last full synchronization.
	lastFullSync time.Time
	// requireFullSync indicates that client state has changed that a full sync is warranted.
	requireFullSync bool
	resourceTypes   map[ResourceType]bool
	http            *http.Client
	syncToken       string
	apiToken        string
	State           *State
}

func NewSyncClient(apiToken string) *SyncClient {
	return &SyncClient{
		http:          &http.Client{},
		apiToken:      apiToken,
		syncToken:     "*",
		resourceTypes: map[ResourceType]bool{},
		State:         NewState(),
	}
}

// UseResources marks the resource types to be synchronized.
func (sc *SyncClient) UseResources(resourceTypes ...ResourceType) {
	for _, resourceType := range resourceTypes {
		if resourceType != Items {
			// Log a warning or handle the case where the resource type is not implemented
			fmt.Printf("WARN : Resource type %v not implemented\n", resourceType)
			continue
		}

		if sc.resourceTypes[resourceType] == false {
			sc.resourceTypes[resourceType] = true

			// Incremental sync may not contain all necessary data, so require a full sync.
			sc.requireFullSync = true
		}
	}
}

// headers applies common headers to the given request.
func (sc *SyncClient) headers(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+sc.apiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)
}

// get performs a GET request to the Todoist API, building a request with the given path and parameters.
// It will also apply Authorization, Content-Type, Accept, and User-Agent headers.
func (sc *SyncClient) get(path string, params url.Values) (*http.Response, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	url := API_BASE_URL + path
	if encodedParams := params.Encode(); len(encodedParams) > 0 {
		url += "?" + encodedParams
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	sc.headers(req)

	fmt.Printf("DEBUG : Sending GET request to %s\n", req.URL.String())
	return sc.http.Do(req)
}

// post performs a POST request to the Todoist API, building a request with the given path and parameters.
// It will also apply Authorization, Content-Type, Accept, and User-Agent headers.
func (sc *SyncClient) post(path string, params url.Values, body []byte) (*http.Response, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	url := API_BASE_URL + path
	if encodedParams := params.Encode(); len(encodedParams) > 0 {
		url += "?" + encodedParams
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	sc.headers(req)

	fmt.Printf("DEBUG : Sending POST request to %s with body: %s\n", req.URL.String(), string(body))

	res, err := sc.http.Do(req)

	fmt.Printf("DEBUG : Received response with status code %d\n", res.StatusCode)

	return res, err
}
