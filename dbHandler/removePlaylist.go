package dbHandler

import (
	"fmt"
)

func (db *Database) RemovePlaylist(playlistID string) error {
	// TODO enforce foreign key constraints
	// TODO remove all songs in the playlist
	query := ("DELETE FROM playlists WHERE DIR_ID = ?")
	_, err := db.db.Exec(query, playlistID)
    if err != nil {
        return fmt.Errorf("error removing playlist from db: %w", err)
    }
    return nil
}