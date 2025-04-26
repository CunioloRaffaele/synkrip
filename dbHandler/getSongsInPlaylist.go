package dbHandler

import (
	"log"
)

func (db *Database) GetSongs(dir string) ([]Song, error) {
query := `SELECT * FROM songs WHERE DIR_ID = ?`
	rows, err := db.db.Query(query, dir)
    if err != nil {
        log.Printf("Error executing query: %v", err)
        return nil, err
    }
	defer rows.Close()
	songs := []Song{}
	for rows.Next() {
		var song Song
		err := rows.Scan(&song.SONG_NAME, &song.SONG_ALBUM_NAME, &song.SONG_ARTIST_NAME, &song.SONG_YT_ID,&song.SONG_PLATFORM_ID, &song.IS_DOWNLOADED,&song.DIR_ID, &song.ADD_DATE)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}