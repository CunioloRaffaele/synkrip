package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// needs spotify api url not public url
// Example URL: https://api.spotify.com/v1/playlists/0NgKG9cRZ96xjsTxQHCVCY/tracks
func getSongsInPlaylist(url string) playlistsStruct {
    log.Println("spotify Fetching from:", url)
    method := "GET"
    client := &http.Client{}
    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        log.Println("Error creating request:", err)
        return playlistsStruct{}
    }

    // Get the token
    token, err := getToken()
    if err != nil {
        log.Println("Error getting token:", err)
        return playlistsStruct{}
    }
    req.Header.Add("Authorization", "Bearer "+token)

    // Execute the request
    res, err := client.Do(req)
    if err != nil {
        log.Println("Error executing request:", err)
        return playlistsStruct{}
    }
    defer res.Body.Close()

    // Read the response body
    body, err := io.ReadAll(res.Body)
    if err != nil {
        log.Println("Error reading response body:", err)
        return playlistsStruct{}
    }

    // Check for non-2xx status codes
    if res.StatusCode < 200 || res.StatusCode >= 300 {
        log.Printf("Error: received status code %d\n", res.StatusCode)
        return playlistsStruct{}
    }

    // Unmarshal the response into the playlistsStruct
    var playlists playlistsStruct
    if err := json.Unmarshal(body, &playlists); err != nil {
        log.Println("Error unmarshaling response:", err)
        return playlistsStruct{}
    }

    // If there is a "next" URL, recursively fetch the next batch of songs
    if playlists.Next != "" && playlists.Next != "null" {
        log.Println("Fetching next batch of songs from:", playlists.Next)
        nextBatch := getSongsInPlaylist(playlists.Next)
        playlists.Items = append(playlists.Items, nextBatch.Items...) // Append the next batch of songs
    }

    return playlists
}

// need spotify public url not api url
// Example URL: https://open.spotify.com/playlist/0NgKG9cRZ96xjsTxQHCVCY?si=75cda28f62e14d86
func getPlaylistId (url string) string {
    // Split the URL by "playlist/" to isolate the playlist ID
    parts := strings.Split(url, "playlist/")
    if len(parts) < 2 {
        log.Println("Invalid Spotify playlist URL format")
        return ""
    }
    // Split the second part by "?" to remove any query parameters
    idParts := strings.Split(parts[1], "?")
    return idParts[0]
}

// needs spotify api url not public url
// Example URL: https://api.spotify.com/v1/playlists/0NgKG9cRZ96xjsTxQHCVCY/tracks
func getPlaylistInfo (url string) playlistsStructFull {
    log.Println("spotify Fetching from:", url)
    method := "GET"
    client := &http.Client{}
    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        log.Println("Error creating request:", err)
        return playlistsStructFull{}
    }

    // Get the token
    token, err := getToken()
    if err != nil {
        log.Println("Error getting token:", err)
        return playlistsStructFull{}
    }
    req.Header.Add("Authorization", "Bearer "+token)

    // Execute the request
    res, err := client.Do(req)
    if err != nil {
        log.Println("Error executing request:", err)
        return playlistsStructFull{}
    }
    defer res.Body.Close()

    // Read the response body
    body, err := io.ReadAll(res.Body)
    if err != nil {
        log.Println("Error reading response body:", err)
        return playlistsStructFull{}
    }

    // Check for non-2xx status codes
    if res.StatusCode < 200 || res.StatusCode >= 300 {
        log.Printf("Error: received status code %d\n", res.StatusCode)
        return playlistsStructFull{}
    }

    // Unmarshal the response into the playlistsStruct
    var playlists playlistsStructFull
    if err := json.Unmarshal(body, &playlists); err != nil {
        log.Println("Error unmarshaling response:", err)
        return playlistsStructFull{}
    }

    return playlists
}

func imageToBase64 (playlistImage string) string {
    // Convert the image URL to base64
    if playlistImage != "" {
        resp, err := http.Get(playlistImage)
        if err != nil {
            log.Println("Error fetching playlist image:", err)
        } else {
            defer resp.Body.Close()
            imageData, err := io.ReadAll(resp.Body)
            if err != nil {
                log.Println("Error reading image data:", err)
            } else {
                playlistImage = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imageData)
            }
        }
    }
    return playlistImage
}

// need spotify public url not api url
// Example URL: https://open.spotify.com/playlist/0NgKG9cRZ96xjsTxQHCVCY?si=75cda28f62e14d86
func IngestSpotifyPlaylist (url string) (playlist playlistsStruct, playlistName string, playlistImage string) {
    playlistId := getPlaylistId(url)
    apiUrlTraks := fmt.Sprintf(playlistEndpointTracks, playlistId)
    apiUrlName := fmt.Sprintf(playlistEndpointName, playlistId)
    pName := getPlaylistInfo(apiUrlName).Name
    image := imageToBase64(getPlaylistInfo(apiUrlName).Images[0].URL)
    return getSongsInPlaylist(apiUrlTraks), pName, image
}

//TODO: handle error properly