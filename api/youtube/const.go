package youtube

// Base URLs
const (
    baseUrl = "https://music.youtube.com"
)

// API Endpoints
const (
    searchEndpoint = baseUrl + "/youtubei/v1/search"
)

// Payload for Search (needs string with title)
const (
	searchPayload = `{
    "context": {
        "client": {
            "clientName": "WEB_REMIX",
            "clientVersion": "1.20250317.01.00"
        }
    },
    "query": "%s"
}`
)