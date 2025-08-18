package main

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
	fs "synkrip/fsHandler"

	"github.com/google/uuid"
)

// Settings structure to hold user data
type Settings struct {
	EulaAccepted      bool     `json:"eulaAccepted"`
	TelemetryEnabled  bool    `json:"telemetryEnabled"`
	RecentLibraries   []string `json:"recentLibraries"`
	UserIdentifier    string   `json:"userIdentifier"`
	UpdateCheck       bool     `json:"updateCheck"`
}

func NewDefaultSettings() *Settings {
	userID, _ := getUserIdentifier()
	return &Settings{
		EulaAccepted:    false,
		TelemetryEnabled:true,
		RecentLibraries: []string{},
		UserIdentifier:  userID,
		UpdateCheck:     true,
	}
}

func getUserIdentifier() (string, error) {
	currentUser, err := user.Current()
    if err != nil {
        return "", err
    }
    
    identifier := currentUser.Username
	userUUID := uuid.New()

    userIdentifier := identifier + "-" + userUUID.String()

	return userIdentifier, nil
}

// GetSettings loads settings from the settings file
func GetSettings() (*Settings, error) {
	appDataPath, err := fs.GetAppDataPath()
	if err != nil {
		return nil, err
	}

	settingsFilePath := appDataPath + "/settings.json"

	// Check if the file exists
	if _, err := os.Stat(settingsFilePath); os.IsNotExist(err) {
		// create a new settings file with default values
		defaultSettings := NewDefaultSettings()
		file, err := os.Create(settingsFilePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		err = json.NewEncoder(file).Encode(defaultSettings)
		if err != nil {
			return nil, err
		}
		log.Println("Settings file not found, creating a new one with default values.")
		log.Print(defaultSettings)
		return defaultSettings, nil
	}

	// Open the file
	file, err := os.Open(settingsFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the JSON data
	var settings Settings
	err = json.NewDecoder(file).Decode(&settings)
	if err != nil {
		return nil, err
	}

	log.Println("Settings file found, loading values.")
	log.Print(settings)
	return &settings, nil
}

func SaveSettings(settings *Settings) error {
    appDataPath, err := fs.GetAppDataPath()
    if err != nil {
        return err
    }

    settingsFilePath := appDataPath + "/settings.json"

    // Ensure the directory exists
    err = os.MkdirAll(appDataPath, os.ModePerm)
    if err != nil {
        return err
    }

    // Create or overwrite the file
    file, err := os.Create(settingsFilePath)
    if err != nil {
        return err
    }
    defer file.Close()

    // Encode the settings as JSON
    err = json.NewEncoder(file).Encode(settings)
    if err != nil {
        return err
    }

    log.Println("Settings saved successfully.")
    return nil
}

func (s *Settings) AddLibrary(path string) {
	if len(path) == 0 {
		return
	}
	// Check if the path already exists in the recent libraries
	for _, lib := range s.RecentLibraries {
		if lib == path {
			return // Path already exists, no need to add it again
		}
	}
	s.RecentLibraries = append(s.RecentLibraries, path)
	if (len(s.RecentLibraries) > 3) {
		s.RecentLibraries = s.RecentLibraries[1:] // Keep only the last 3 libraries
	}
}
