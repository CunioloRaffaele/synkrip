package youtube

// Base URLs
const (
    baseUrl = "https://music.youtube.com"
    privateApiUrl = "https://www.youtube.com/youtubei/v1"
)

// API Endpoints
const (
    searchEndpoint = baseUrl + "/youtubei/v1/search"
    playlistEndpoint = privateApiUrl + "/browse"
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

// Payload for Playlist
const (
	playlistPayload = `{
    "context": {
        "client": {
            "clientName": "WEB",
            "clientVersion": "2.20250502.01.00"
        }
    },
    "browseId": "VL%s"
}`
)