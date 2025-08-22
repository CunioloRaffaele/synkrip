package dbHandler

import (
    "database/sql"
    "encoding/json"
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
                ADD_DATE:         addDateSong.String,
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