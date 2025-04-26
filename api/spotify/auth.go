package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var accessTokenStruct authResponse // Cached token struct

// getBearer fetches a new token from Spotify
// Do NOT use this function directly, use getToken() instead
func getBearer() (authResponse, error) {
  method := "POST"
  payload := strings.NewReader("grant_type=client_credentials")

  client := &http.Client{}
  req, err := http.NewRequest(method, authUrl, payload)
  if err != nil {
    log.Println(err)
    return authResponse{}, err
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Authorization", "Basic " + authApiBasic)

  res, err := client.Do(req)
  if err != nil {
    return authResponse{}, err
  }
  defer res.Body.Close()

  if res.StatusCode < 200 || res.StatusCode >= 300 {
    return authResponse{}, fmt.Errorf("failed to get token from Spotify Api, status code: %d", res.StatusCode)
  }

  body, err := io.ReadAll(res.Body)
  if err != nil {
    return authResponse{}, err
  }

  var authResponseS authResponse
  if err := json.Unmarshal(body, &authResponseS); err != nil {
    return authResponse{}, err
  }

  // Add the current time to calculate expiration
  authResponseS.ExpiresTime = time.Now().UTC().Unix() + int64(authResponseS.ExpiresIn) - 20

  log.Println("Got Spotify API auth response:", authResponseS)
  return authResponseS, nil
}

// getToken ensures a valid token is available (fetches a new one if needed)
func getToken() (string, error) {
	// Check if the token is already initialized and valid
	if accessTokenStruct.AccessToken != "" || time.Now().UTC().Unix() < accessTokenStruct.ExpiresTime {
		log.Println("Using cached token for Spotify API")
		return accessTokenStruct.AccessToken, nil
	}

	// Fetch a new token if not initialized or expired
	log.Println("Fetching a new token for Spotify API")
	newToken, err := getBearer()
	if err != nil {
		return "", err
	}

	// Cache the new token
	accessTokenStruct = newToken
	return accessTokenStruct.AccessToken, nil
}