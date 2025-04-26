package dbHandler

func (db *Database) RemoveSongFromPlaylist(playlistID string, songID string) error {
	stmt, err := db.db.Prepare("DELETE FROM songs WHERE DIR_ID = ? AND SONG_PLATFORM_ID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement with the provided playlist ID and song ID
	_, err = stmt.Exec(playlistID, songID)
	if err != nil {
		return err
	}

	return nil
}