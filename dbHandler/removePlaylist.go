package dbHandler

import (
	"fmt"
)

func (db *Database) RemovePlaylist(playlistID string) error {
	queryRemovePlaylist := ("DELETE FROM playlists WHERE DIR_ID = ?")
	_, err := db.db.Exec(queryRemovePlaylist, playlistID)
    if err != nil {
        return fmt.Errorf("error removing playlist from db: %w", err)
    }
	queryRemoveSongs := ("DELETE FROM songs WHERE DIR_ID = ?")
	_, err = db.db.Exec(queryRemoveSongs, playlistID)
    if err != nil {
        return fmt.Errorf("error removing songs from db: %w", err)
    }
    return nil
}