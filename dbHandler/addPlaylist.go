package dbHandler

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/text/unicode/norm"
)

// AddPlaylistEntry adds a new entry to the playlist filed of the database
func (db *Database) AddPlaylistEntry(playlistName string, service string, playlistUrl string, image string) error {

	normalizedPlaylistName := norm.NFC.String(playlistName)

	// Check if playlist already exists
	checkQuery := `SELECT COUNT(*) FROM playlists WHERE DIR_ID = ?`
	// Usa la stringa normalizzata per la ricerca
	row := db.db.QueryRow(checkQuery, normalizedPlaylistName)
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println("Error checking if playlist exists:", err)
		return fmt.Errorf("error checking if playlist exists: %w", err)
	}
	if count > 0 {
		log.Println("Playlist already exists in the database")
		// verify that the playlistName is associated with "unknownService"
		var existingService string
		query := `SELECT SERVICE FROM playlists WHERE DIR_ID = ?`
		row := db.db.QueryRow(query, normalizedPlaylistName)
		err := row.Scan(&existingService)
		if err != nil {
			log.Println("Error checking existing playlist service:", err)
			return fmt.Errorf("error checking existing playlist service: %w", err)
		}
		if existingService == "unknownService" {
			log.Println("Updating existing playlist with new service information")
			updateQuery := `UPDATE playlists SET SERVICE = ?, URL = ?, IMAGE = ?, LAST_MODIFIED = ? WHERE DIR_ID = ?`
			// E anche qui
			_, err := db.db.Exec(updateQuery, service, playlistUrl, image, time.Now(), normalizedPlaylistName)
			if err != nil {
				log.Println("Error updating existing playlist:", err)
				return fmt.Errorf("error updating existing playlist: %w", err)
			}
			return nil
		}
		// Se il servizio non è "unknownService", considerala un duplicato e non fare nulla.
		return fmt.Errorf("playlist '%s' already exists with service '%s'", normalizedPlaylistName, existingService)
	}

	// Insert new entry into the database
	query := `INSERT INTO playlists (DIR_ID, SERVICE, URL, IMAGE, ADD_DATE, LAST_MODIFIED) VALUES (?, ?, ?, ?, ?, ?)`
	// Salva sempre la stringa normalizzata
	_, err = db.db.Exec(query, normalizedPlaylistName, service, playlistUrl, image, time.Now(), time.Now())
	if err != nil {
		log.Println("Error inserting new playlist entry:", err)
		return fmt.Errorf("error inserting new playlist entry: %w", err)
	}
	log.Println("New playlist entry added to the database")
	return nil
}