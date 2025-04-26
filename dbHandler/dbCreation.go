package dbHandler

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateDatabase creates a SQLite database at the specified path
func CreateDatabase(path string) (error){
	// Ensure directory exists
	if err := os.MkdirAll(path, 0755); err != nil {
		fmt.Println("failed to create directory: %w", err)
		return err
	}
	dbPath := filepath.Join(path, ".synkrip.db")
	// Verify if the database already exists
	if _, err := os.Stat(dbPath); err == nil {
		fmt.Println("Database already exists at:", dbPath)
		return fmt.Errorf("Database already exists at: %s", dbPath)
	}
	// create the database called synkrip.db
	db, err := os.Create(dbPath)
	if err != nil {
		fmt.Println("failed to create database: %w", err)
		return err
	}
	
	defer db.Close()
	return nil
}