package dbHandler

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
)

func (db *Database) GetDBJson() (string, error) {
    // Map to group songs by playlist
    playlistsMap := make(map[string]*Playlist)

    // Execute the query
    query := `SELECT 
              playlists.DIR_ID, playlists.SERVICE, playlists.URL, playlists.IMAGE, playlists.ADD_DATE, playlists.LAST_MODIFIED,
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
        var dirID, service, url, image, addDatePlaylist, lastModified string
        var songName, songAlbumName, songArtistName, songYTID, addDateSong sql.NullString
        var isDownloaded sql.NullInt64

        // Scan the row
        err := rows.Scan(&dirID, &service, &url, &image, &addDatePlaylist, &lastModified, &songName, &songAlbumName, &songArtistName, &songYTID, &isDownloaded, &addDateSong)
        if err != nil {
            log.Printf("Error scanning row: %v", err)
            return "", err
        }

        // Check if the playlist already exists in the map
        playlist, exists := playlistsMap[dirID]
        if !exists {
            // Create a new playlist if it doesn't exist
            playlist = &Playlist{
                DIR_ID:        dirID,
                SERVICE:       service,
                URL:           url,
                IMAGE:         image,
                ADD_DATE:      addDatePlaylist,
                LAST_MODIFIED: lastModified,
                Songs:         []Song{},
                ToBeSynced:    false,
                DuplicateCount: 0,
                // Inizializza il set per il controllo dei duplicati
                _songSet: make(map[string]struct{}),
            }
            playlistsMap[dirID] = playlist
        }

        // Add the song to the playlist if it exists
        if songName.Valid {

            isDuplicateFlag := 0
            songKey := fmt.Sprintf("%s||%s", songName.String, songArtistName.String)

            // Check if the song has already been seen in this playlist
            if _, found := playlist._songSet[songKey]; found {
                isDuplicateFlag = 1
                playlist.DuplicateCount++
                //log.Printf("Duplicate song flagged in playlist '%s': %s by %s", dirID, songName.String, songArtistName.String)
            } else {
                // add to the set for future checks
                playlist._songSet[songKey] = struct{}{}
            }

            song := Song{
                SONG_NAME:        songName.String,
                SONG_ALBUM_NAME:  songAlbumName.String,
                SONG_ARTIST_NAME: songArtistName.String,
                SONG_YT_ID:       songYTID.String,
                IS_DOWNLOADED:    int(isDownloaded.Int64),
                ADD_DATE:         addDateSong.String,
                Duplicate:        isDuplicateFlag, // Imposta il flag
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