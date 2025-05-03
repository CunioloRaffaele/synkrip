package dbHandler

import (
	"fmt"
	"log"
)

func (db *Database) UpdateImage(playlistID string, image string) error {
    query := "UPDATE playlists SET image = ? WHERE DIR_ID = ?"
    _, err := db.db.Exec(query, image, playlistID)
    if err != nil {
		log .Println("Error updating playlist image:", err)
        return fmt.Errorf("failed update playlist image: %w", err)
    }
    return nil
}