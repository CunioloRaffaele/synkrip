package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	fsHandler "synkrip/fsHandler"
	storage_go "github.com/supabase-community/storage-go"
)

// TelemetryConfig represents the structure of telemetry.json
type TelemetryConfig struct {
	Supabase struct {
		URL     string `json:"url"`
		AnonKey string `json:"anon_key"`
	} `json:"supabase"`
}

// loadTelemetryConfig reads the telemetry.json file
func loadTelemetryConfig() (*TelemetryConfig, error) {
	configFile := "telemetry.json"

	// Check if file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, err
	}

	// Read the file
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	// Parse JSON
	var config TelemetryConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
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

	// Read the file content
	fileBody, err := os.ReadFile(logFilePath)
	if err != nil {
		log.Printf("Error reading log file: %v", err)
		return
	}

	upsert := true // if file exists, update it
	resultUPLOAD, err := storageClient.UploadFile("Logs", a.Settings.UserIdentifier+".txt", bytes.NewReader(fileBody), storage_go.FileOptions{Upsert: &upsert})
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("Telemetry data uploaded: %+v", resultUPLOAD)
}