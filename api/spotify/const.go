package spotify

// Base URLs
const (
    baseUrl     = "https://api.spotify.com/v1"
    authUrl     = "https://accounts.spotify.com/api/token"
    authApiBasic= "NTMyZDNlMTczNDc2NDU2YWI2NGYyZWJkMjk3OGE0ZDQ6Y2I5MTdjNDJlNDBlNGM0MWJhN2EyYzQ1Mjg1MzUyOTQ="
)

// API Endpoints
const (
    playlistEndpointTracks = baseUrl + "/playlists/%s/tracks"   // needs playlist_id
    playlistEndpointName   = baseUrl + "/playlists/%s"          // needs playlist_id
    trackEndpoint          = baseUrl + "/tracks/"               // needs track_id  
)