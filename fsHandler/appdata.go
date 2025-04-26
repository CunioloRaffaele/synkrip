package fsHandler

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetAppDataPath() (string, error) {
    var appDataPath string

    if runtime.GOOS == "darwin" {
        // macOS: ~/Library/Application Support/<YourAppName>
        homeDir, err := os.UserHomeDir()
        if err != nil {
            return "", err
        }
        appDataPath = filepath.Join(homeDir, "Library", "Application Support", "SynkRip")
    } else if runtime.GOOS == "windows" {
        // Windows: %AppData%\<YourAppName>
        appDataDir := os.Getenv("APPDATA")
        if appDataDir == "" {
            return "", fmt.Errorf("APPDATA environment variable is not set")
        }
        appDataPath = filepath.Join(appDataDir, "SynkRip")
	}
	
    // Ensure the directory exists
    err := os.MkdirAll(appDataPath, 0755)
    if err != nil {
        return "", err
    }

    return appDataPath, nil
}