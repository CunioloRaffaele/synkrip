package dbHandler

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// Database represents a connection to the SQLite database
type Database struct {
    db *sql.DB
}

// Connect establishes a connection to the database at the given path
// If the database file doesn't exist, it will be created
func Connect(dbPath string) (*Database, error) {
	dbPath = filepath.Join(dbPath, ".synkrip.db")
	// Check if database file exists before attempting to connect
    _, err := os.Stat(dbPath)
    if os.IsNotExist(err) {
        return nil, fmt.Errorf("database does not exist at %s", dbPath)
    } else if err != nil {
        return nil, fmt.Errorf("error checking database file: %w", err)
    }

	// Open the database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return nil, err
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Printf("Failed to ping database: %v", err)
		db.Close()
		return nil, err
	} else {
		log.Printf("Connected to database at %s", dbPath)
	}

	return &Database{db: db}, nil
}

// Close closes the database connection
func (db *Database) Close() error {
	if db.db != nil {
		return db.db.Close()
	}
	return nil
}
	