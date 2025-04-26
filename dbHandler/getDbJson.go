package dbHandler

import (
    "database/sql"
    "encoding/json"
    "log"
)

func (db *Database) GetDBJson() (string, error) {
    // Define a structure to hold the JSON data
    /*type Song struct {
        SONG_NAME        string `json:"song_name"`
        SONG_ALBUM_NAME  string `json:"song_album_name"`
        SONG_ARTIST_NAME string `json:"song_artist_name"`
        SONG_YT_ID       string `json:"song_yt_id"`
        IS_DOWNLOADED    int    `json:"is_downloaded"`
        ADD_DATE         string `json:"add_date"`
    }

    type Playlist struct {
        DIR_ID   string `json:"dir_id"`
        SERVICE  string `json:"service"`
        URL      string `json:"url"`
        IMAGE    string `json:"image"`
        Songs    []Song `json:"songs"`
    }*/

    // Map to group songs by playlist
    playlistsMap := make(map[string]*Playlist)

    // Execute the query
    query := `SELECT 
              playlists.DIR_ID, playlists.SERVICE, playlists.URL, playlists.IMAGE,
              songs.SONG_NAME, songs.SONG_ALBUM_NAME, songs.SONG_ARTIST_NAME, songs.SONG_YT_ID, songs.IS_DOWNLOADED, songs.ADD_DATE
              FROM playlists LEFT JOIN songs ON playlists.DIR_ID = songs.DIR_ID`
	rows, err := db.db.Query(query)
    if err != nil {
        log.Printf("Error executing query: %v", err)
        return "", err
    }
    defer rows.Close()

    // Process the rows
    for rows.Next() {
        var dirID, service, url, image string
        var songName, songAlbumName, songArtistName, songYTID, addDate sql.NullString
        var isDownloaded sql.NullInt64

        // Scan the row
        err := rows.Scan(&dirID, &service, &url, &image, &songName, &songAlbumName, &songArtistName, &songYTID, &isDownloaded, &addDate)
        if err != nil {
            log.Printf("Error scanning row: %v", err)
            return "", err
        }

        // Check if the playlist already exists in the map
        playlist, exists := playlistsMap[dirID]
        if !exists {
            // Create a new playlist if it doesn't exist
            playlist = &Playlist{
                DIR_ID:  dirID,
                SERVICE: service,
                URL:     url,
                IMAGE:   image,
                Songs:   []Song{},
                ToBeSynced: false,
            }
            playlistsMap[dirID] = playlist
        }

        // Add the song to the playlist if it exists
        if songName.Valid {
            song := Song{
                SONG_NAME:        songName.String,
                SONG_ALBUM_NAME:  songAlbumName.String,
                SONG_ARTIST_NAME: songArtistName.String,
                SONG_YT_ID:       songYTID.String,
                IS_DOWNLOADED:    int(isDownloaded.Int64),
                ADD_DATE:         addDate.String,
            }
            if song.IS_DOWNLOADED == 0 {
                playlist.ToBeSynced = true
            }
            playlist.Songs = append(playlist.Songs, song)
        }
    }

    // Convert the map to a slice of playlists
    playlists := make([]Playlist, 0, len(playlistsMap))
    for _, playlist := range playlistsMap {
        playlists = append(playlists, *playlist)
    }

    // Convert the playlists slice to JSON
    jsonData, err := json.MarshalIndent(playlists, "", "  ")
    if err != nil {
        log.Printf("Error marshaling playlists to JSON: %v", err)
        return "", err
    }

    return string(jsonData), nil
}