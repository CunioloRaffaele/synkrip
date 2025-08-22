package dbHandler

import (
	"log"
)

// First time setup of the database
func (db *Database) Setup() error {
	initStatements := `
	CREATE TABLE playlists (
		DIR_ID TEXT PRIMARY KEY,
		SERVICE TEXT NOT NULL,
		URL TEXT NOT NULL,
		IMAGE TEXT NOT NULL,
		ADD_DATE DATETIME NOT NULL,
		LAST_MODIFIED DATETIME NOT NULL
	);

	CREATE TABLE songs (
		SONG_NAME TEXT NOT NULL,
		SONG_ALBUM_NAME TEXT NOT NULL,
		SONG_ARTIST_NAME TEXT NOT NULL,
		SONG_YT_ID TEXT NOT NULL,
		SONG_PLATFORM_ID TEXT NOT NULL,
		IS_DOWNLOADED INTEGER NOT NULL,
		DIR_ID TEXT NOT NULL,
		ADD_DATE TEXT NOT NULL,
		FOREIGN KEY (DIR_ID) REFERENCES playlists(DIR_ID) ON DELETE CASCADE
	);
	`
	_, err := db.db.Exec(initStatements)
	if err != nil {
		log.Printf("Failed to execute setup statements: %v", err)
		return err
	}
	log.Println("Database setup completed successfully.")
	return nil
}