package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"

	fsHandler "synkrip/fsHandler"
	storage_go "github.com/supabase-community/storage-go"
)

const MaxLogSize = 1024 * 1024 // 1MB in bytes

// TelemetryConfig represents the structure of telemetry.json
type TelemetryConfig struct {
	Supabase struct {
		URL     string `json:"url"`
		AnonKey string `json:"anon_key"`
	} `json:"supabase"`
}

// Build-time variables (set by ldflags)
var (
    SupabaseURL     string
    SupabaseAnonKey string
)

// loadTelemetryConfig loads config from build-time variables or returns disabled config
func loadTelemetryConfig() (*TelemetryConfig, error) {
    // Crash if build-time variables are empty
    if SupabaseURL == "" || SupabaseAnonKey == "" {
        log.Println("Telemetry disabled: no build-time configuration. Killing...")
        os.Exit(1)
    }

    config := &TelemetryConfig{}
    config.Supabase.URL = SupabaseURL
    config.Supabase.AnonKey = SupabaseAnonKey
    
    return config, nil
}

// readLastMBOfFile reads the last 1MB of a file, or the entire file if it's smaller
func readLastMBOfFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()

	// If file is smaller than 1MB, read the entire file
	if fileSize <= MaxLogSize {
		log.Printf("Log file size: %d bytes (less than 1MB), reading entire file", fileSize)
		return os.ReadFile(filePath)
	}

	// File is larger than 1MB, read only the last 1MB
	log.Printf("Log file size: %d bytes (larger than 1MB), reading last 1MB", fileSize)

	// Calculate the offset to start reading from (last 1MB)
	offset := fileSize - MaxLogSize

	// Seek to the calculated offset
	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Read the last 1MB
	buffer := make([]byte, MaxLogSize)
	bytesRead, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, err
	}

	// Return only the bytes that were actually read
	return buffer[:bytesRead], nil
}

func uploadTelemetry(a *App) {
	if !a.Settings.TelemetryEnabled {
		log.Println("Telemetry is disabled in settings, skipping upload.")
		return
	}

	// Load configuration from JSON file
	config, err := loadTelemetryConfig()
	if err != nil {
		log.Printf("Error loading telemetry config: %v. Skipping telemetry upload.", err)
		return
	}

	// Create storage client using config
	storageClient := storage_go.NewClient(config.Supabase.URL, config.Supabase.AnonKey, nil)

	appDataPath, _ := fsHandler.GetAppDataPath()
	logFilePath := filepath.Join(appDataPath, "log.txt")

	// Check if file exists
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		log.Printf("Log file does not exist: %s", logFilePath)
		return
	}

	// Read the file content (last 1MB if file is large)
	fileBody, err := readLastMBOfFile(logFilePath)
	if err != nil {
		log.Printf("Error reading log file: %v", err)
		return
	}

	log.Printf("Uploading %d bytes of log data", len(fileBody))

	upsert := true // if file exists, update it
	resultUPLOAD, err := storageClient.UploadFile("Logs", a.Settings.UserIdentifier+".txt", bytes.NewReader(fileBody), storage_go.FileOptions{Upsert: &upsert})
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("Telemetry data uploaded: %+v", resultUPLOAD)
}