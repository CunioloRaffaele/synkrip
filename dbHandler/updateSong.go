package dbHandler

import (
	"fmt"
	"log"
)

func (db *Database) MarkSongAsDownloaded(playlist string, ytId string) error {
    query := "UPDATE songs SET is_downloaded = 1 WHERE DIR_ID = ? AND SONG_YT_ID = ?"
    _, err := db.db.Exec(query, playlist, ytId)
    if err != nil {
		log .Println("Error marking song as downloaded:", err)
        return fmt.Errorf("failed to mark song as downloaded: %w", err)
    }
    return nil
}