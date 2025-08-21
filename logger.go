package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"synkrip/fsHandler"
)

var logFile *os.File // Variabile globale per poter chiudere il file alla fine

func loggerSetup() {
	path, err := fsHandler.GetAppDataPath()
	if err != nil {
		log.Printf("Error getting app data path: %v", err)
		return
	}

	// Use filepath.Join instead of string concatenation
	logPath := filepath.Join(path, "log.txt")

	// Create log file if it doesn't exist
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		txtFile, err := os.Create(logPath)
		if err != nil {
			log.Panic(err)
		}
		txtFile.Close()
	}

	// Open log file with proper flags
	logFile, err = os.OpenFile(logPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}

	// Create a MultiWriter that writes to both stdout and file
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Set log output to both terminal and file
	log.SetOutput(multiWriter)

	// Set log flags for better debugging
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("--------------------START-----------------------")
	log.Println("Starting program with logger set to custom file")
	log.Printf("Log file location: %s", logPath)

	// Force flush to ensure the message is written immediately
	logFile.Sync()
}

// Funzione per chiudere il file di log
func loggerCleanup() {
	if logFile != nil {
		log.Println("--------------------END-----------------------")
		log.Println("Closing logger and program")

		// Force flush before closing
		logFile.Sync()
		logFile.Close()
	}
}