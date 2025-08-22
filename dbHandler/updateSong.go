package dbHandler

import (
	"fmt"
	"log"
	"time"
)

func (db *Database) MarkSongAsDownloaded(playlist string, ytId string) error {
    // update song table
    query := "UPDATE songs SET is_downloaded = 1 WHERE DIR_ID = ? AND SONG_YT_ID = ?"
    _, err := db.db.Exec(query, playlist, ytId)
    if err != nil {
		log .Println("Error marking song as downloaded:", err)
        return fmt.Errorf("failed to mark song as downloaded: %w", err)
    }
    // update playlist table to renew last modified date
    query = "UPDATE playlists SET LAST_MODIFIED = ? WHERE DIR_ID = ?"
    _, err = db.db.Exec(query, time.Now(), playlist)
    if err != nil {
        log.Println("Error updating playlist last modified date:", err)
        return fmt.Errorf("failed to update playlist last modified date: %w", err)
    }
    return nil
}