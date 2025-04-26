package spotify

import "time"

// Generated with https://transform.tools/json-to-go
type playlistsStruct struct {
	Href  string `json:"href"`
	Items []struct {
		AddedAt time.Time `json:"added_at"`
		AddedBy struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"added_by"`
		IsLocal      bool        `json:"is_local"`
		PrimaryColor string `json:"primary_color"`
		Track        struct {
			PreviewURL       string `json:"preview_url"`
			AvailableMarkets []string    `json:"available_markets"`
			Explicit         bool        `json:"explicit"`
			Type             string      `json:"type"`
			Episode          bool        `json:"episode"`
			Track            bool        `json:"track"`
			Album            struct {
				AvailableMarkets []string `json:"available_markets"`
				Type             string   `json:"type"`
				AlbumType        string   `json:"album_type"`
				Href             string   `json:"href"`
				ID               string   `json:"id"`
				Images           []struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"images"`
				Name                 string `json:"name"`
				ReleaseDate          string `json:"release_date"`
				ReleaseDatePrecision string `json:"release_date_precision"`
				URI                  string `json:"uri"`
				Artists              []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"artists"`
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				TotalTracks int `json:"total_tracks"`
			} `json:"album"`
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			DiscNumber  int `json:"disc_number"`
			TrackNumber int `json:"track_number"`
			DurationMs  int `json:"duration_ms"`
			ExternalIds struct {
				Isrc string `json:"isrc"`
			} `json:"external_ids"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href       string `json:"href"`
			ID         string `json:"id"`
			Name       string `json:"name"`
			Popularity int    `json:"popularity"`
			URI        string `json:"uri"`
			IsLocal    bool   `json:"is_local"`
		} `json:"track"`
		VideoThumbnail struct {
			URL string `json:"url"`
		} `json:"video_thumbnail"`
	} `json:"items"`
	Limit    int         `json:"limit"`
	Next     string `json:"next"`
	Offset   int         `json:"offset"`
	Previous string `json:"previous"`
	Total    int         `json:"total"`
}