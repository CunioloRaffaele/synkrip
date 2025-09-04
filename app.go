package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"synkrip/api"
	"synkrip/dbHandler"
	"synkrip/fsHandler"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context								// Context
	LibPath    string   							// Path to music library
	CurrentDB *dbHandler.Database					// Current database handler
	CurrentFileSystem *fsHandler.FileSystem			// Current file system struct
	Settings *Settings								// Settings struct
	cancelFunc context.CancelFunc					// Cancel function for download
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.CurrentFileSystem = &fsHandler.FileSystem{}
	loggerSetup()
	var err error
	a.Settings, err = GetSettings()
	if err != nil {
		log.Println("Error loading settings at startup:", err)
		return
	}
	checkForUpdate(a, false)
	externalFrameworksInit()
	//spotify.SpotifyTets()
	//youtube.YoutubeTest()
	//uploadTelemetry(a)
}

func (a *App) domready(ctx context.Context) {
	if !a.Settings.EulaAccepted {
		rt.EventsEmit(ctx, "showEula")
		a.Settings.EulaAccepted = true
	}
}

func (a *App) shutdown(ctx context.Context) {
	// close db connection
    if a.CurrentDB != nil {
        a.CurrentDB.Close()
        a.CurrentDB = nil
    }
	// save modified settings to file
	if a.Settings != nil {
		if err := SaveSettings(a.Settings); err != nil {
			log.Println("Error saving settings:", err)
		}
	}
	loggerCleanup()
	// send telemetry data
	uploadTelemetry(a)
}

// Pass db json to the app frontend
func (a *App) GetDB() (string, error) {
    json, err := a.CurrentDB.GetDBJson()
	return json, err
}

// AddPlaylistEntry
func (a *App) AddEntry(url string, service string) error {
	// TODO make this work in a goroutine

	a.setDownloadStatus("Adding Playlist", true, 1, 1)
	musicService, err := api.GetMusicService(service)
	if err != nil {
		log.Println("Error getting music service:", err)
		a.setDownloadStatus("Adding Playlist", false, 1, 1)
		return err
	}
	playlistData, err := musicService.FetchPlaylist(url)
	if err != nil {
		log.Println("Error fetching playlist:", err)
		a.setDownloadStatus("Adding Playlist", false, 1, 1)
		return err
	}

	// Check if playlist already exists
	for _, playlist := range a.CurrentFileSystem.Directories {
		if playlist.PlaylistName == playlistData.Name {
			log.Println("Folder with playlist name already exist. Modifying database for playlist: ", playlistData.Name)
			rt.MessageDialog(a.ctx, rt.MessageDialogOptions{
				Title:   "Playlist Exists",
				Message: fmt.Sprintf("A folder with the name '%s' already exists locally. You want to modify the sync point for it? Remember the folder name must be the same as the synced playlist", playlistData.Name),
				Type:    "info",
			})
		}
	}

	fsHandler.MkDir(a.LibPath, playlistData.Name)
	err = a.CurrentDB.AddPlaylistEntry(playlistData.Name, service, url, playlistData.Image)
	if (err != nil) {
		log.Println("Error adding entry:", err)
		a.setDownloadStatus("Adding Playlist", false, 1, 1)
		return err
	} else {
		a.setDownloadStatus("Adding Playlist", false, 1, 1)
		err = a.updatePlaylistDb()
	if err != nil {
		log.Println("Failed to update playlist database:", err.Error())
		rt.MessageDialog(a.ctx, rt.MessageDialogOptions{
			Title:   "Update Error",
			Message: fmt.Sprintf("Failed to update playlist database: %v", err),
			Type:    "error",
		})
	}
	}

	return nil
}

func (a *App) DeletePlaylist(title string) error {
	err := a.CurrentDB.RemovePlaylist(title)
	if err != nil {
		log.Println("Error deleting playlist:", err)
		return err
	}
	log.Println("Playlist deleted successfully:", title)
	err = a.updatePlaylistDb()
	if err != nil {
		log.Println("Failed to update playlist database:", err.Error())
		rt.MessageDialog(a.ctx, rt.MessageDialogOptions{
			Title:   "Update Error",
			Message: fmt.Sprintf("Failed to update playlist database: %v", err),
			Type:    "error",
		})
	}
	return nil
}

func (a *App) SyncPlaylist(name string) error {
	// Start the downloadContent function in a Go routine
	// TODO: Handle errors properly

	// Cancel any ongoing download if it exists
    if a.cancelFunc != nil {
        a.cancelFunc()
    }

    // Create a new context with cancellation
    ctx, cancel := context.WithCancel(context.Background())
    a.cancelFunc = cancel
    go func() {
        a.downloadContent(ctx, name)
    }()
    // Return immediately since the operation is running in the background
    return nil
}

func (a *App) StopSync() {
    if a.cancelFunc != nil {
        a.cancelFunc()
        log.Println("Sync process stopped.")
		a.setDownloadStatus("Cancelling...", true, 1, 1)
        a.cancelFunc = nil // Reset the cancel function
    } else {
        log.Println("No active sync to stop.")
    }
}

func (a *App) GetSettings() string {
    // Add a check to prevent a race condition on startup.
    // If settings haven't been loaded into the struct yet, load them now.
    if a.Settings == nil {
        log.Println("Settings were not yet loaded. Loading them on demand.")
        var err error
        a.Settings, err = GetSettings() // This is the function from settings.go
        if err != nil {
            log.Println("Failed to load settings on demand:", err)
            return "{}"
        }
    }

    settingsJSON, err := json.Marshal(a.Settings)
    if err != nil {
        log.Println("Failed to marshal settings:", err)
        return "{}" // Return empty JSON object on error
    }
    return string(settingsJSON)
}

func (a *App) OpenLibrary(newLib string) error {
	err := fsHandler.OpenLibrary(a.ctx, &a.LibPath, &a.CurrentDB, a.CurrentFileSystem, newLib)
	if err != nil {
		log.Println("Failed to open library:", err.Error())
		return err
	}
	err = a.updatePlaylistDb()
	if err != nil {
		log.Println("Failed to update playlist database:", err.Error())
		rt.MessageDialog(a.ctx, rt.MessageDialogOptions{
			Title:   "Update Error",
			Message: fmt.Sprintf("Failed to update playlist database: %v", err),
			Type:    "error",
		})
	}
	return nil
}
