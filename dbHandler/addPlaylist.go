package dbHandler

import (
	"fmt"
	"log"
	"time"
)

// AddPlaylistEntry adds a new entry to the playlist filed of the database
func (db *Database) AddPlaylistEntry(playlistName string, service string, playlistUrl string, image string) error{
	// Check if playlist already exists
	checkQuery := `SELECT COUNT(*) FROM playlists WHERE URL = ?`
	row := db.db.QueryRow(checkQuery, playlistUrl)
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println("Error checking if playlist exists:", err)
		return fmt.Errorf("error checking if playlist exists: %w", err)
	}
	if count > 0 {
		log.Println("Playlist already exists in the database")
		return fmt.Errorf("playlist already exists in the database")
	}

	// Insert new entry into the database
	query := `INSERT INTO playlists (DIR_ID, SERVICE, URL, IMAGE, ADD_DATE, LAST_MODIFIED) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = db.db.Exec(query, playlistName, service, playlistUrl, image, time.Now(), time.Time{})
	if err != nil {
		log.Println("Error inserting new playlist entry:", err)
		return fmt.Errorf("error inserting new playlist entry: %w", err)
	}
	log.Println("New playlist entry added to the database")
	return nil
}