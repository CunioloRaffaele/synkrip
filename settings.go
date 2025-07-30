package main

import (
	"encoding/json"
	"log"
	"os"
	fs "synkrip/fsHandler"
)

// Settings structure to hold user data
type Settings struct {
	RecentLibraries []string `json:"recentLibraries"`
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
		defaultSettings := &Settings{
			RecentLibraries: []string{},
		}
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
	err := SaveSettings(s)
	if err != nil {
		log.Println("Error saving settings:", err)
	}
}

