package dbHandler

import (
	"fmt"
	"log"
	"time"
)

// AddPlaylistEntry adds a new entry to the playlist filed of the database
func (db *Database) AddSongInPlaylist(playlistName string, songName string, albumName string, artistName string, songYTId string, songPlatformId string, isDownloaded bool) error{
	// Check if playlist already exists
	checkQuery := `SELECT COUNT(*) FROM playlists WHERE DIR_ID = ?`
	row := db.db.QueryRow(checkQuery, playlistName)
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println("Error checking if playlist exists:", err)
		return fmt.Errorf("error checking if playlist exists: %w", err)
	}
	if count == 0 {
		log.Println("Playlist not found in the database")
		return fmt.Errorf("playlist not found in the database")
	}
	date := time.Now().Unix()
	// Insert new entry into the database
	query := `INSERT INTO songs (
	SONG_NAME, 
	SONG_ALBUM_NAME, 
	SONG_ARTIST_NAME, 
	SONG_YT_ID, 
	SONG_PLATFORM_ID,
	IS_DOWNLOADED, 
	DIR_ID, 
	ADD_DATE) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = db.db.Exec(query, songName, albumName, artistName, songYTId, songPlatformId, isDownloaded, playlistName, date)
	if err != nil {
		log.Println("Error inserting new song entry:", err)
		return fmt.Errorf("error inserting new song entry: %w", err)
	}
	return nil
}