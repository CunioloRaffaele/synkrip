package dbHandler

import (
	"log"
)


func (db *Database) GetPlaylist() ([]Playlist, error) {
query := `SELECT * FROM playlists`
	rows, err := db.db.Query(query)
    if err != nil {
        log.Printf("Error executing query: %v", err)
        return nil, err
    }
	defer rows.Close()
	playlists := []Playlist{}
	for rows.Next() {
		var playlist Playlist
		err := rows.Scan(&playlist.DIR_ID, &playlist.SERVICE, &playlist.URL, &playlist.IMAGE, &playlist.ADD_DATE, &playlist.LAST_MODIFIED)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}