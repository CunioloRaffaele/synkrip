package main

import (
	"context"
	"encoding/json"
	"log"
	"synkrip/api/spotify"
	"synkrip/api/youtube"
	"synkrip/dbHandler"
	"synkrip/fsHandler"
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
	checkForUpdate(a)
	externalFrameworksInit()
	a.Settings, _ = GetSettings()
	//spotify.SpotifyTets()
	//youtube.YoutubeTest()
}

func (a *App) shutdown(ctx context.Context) {
    if a.CurrentDB != nil {
        a.CurrentDB.Close()
        a.CurrentDB = nil
    }
}

// Pass db json to the app frontend
func (a *App) GetDB() (string, error) {
    json, err := a.CurrentDB.GetDBJson()
	return json, err
}

// AddPlaylistEntry
func (a *App) AddEntry(url string, service string) error {
	// TODO make this work in a goroutine
	// TODO verify if the playlist already exists
	a.setDownloadStatus("Adding Playlist", true, 1, 1)
	switch service {
	case "Spotify":
		playlist, name, image := spotify.IngestSpotifyPlaylist(url)
		fsHandler.MkDir(a.LibPath, name)
		err:= a.CurrentDB.AddPlaylistEntry(name, service, url, image)
		for _, song := range playlist.Items {
			ytId, _ := youtube.GetYTid(song.Track.Name + " " + song.Track.Artists[0].Name)
			err = a.CurrentDB.AddSongInPlaylist(name, song.Track.Name, song.Track.Album.Name, song.Track.Artists[0].Name, ytId, song.Track.ID, false)
		}
		if err != nil {
			log.Println("Error adding entry:", err)
		} else {
			log.Println("Entry added successfully, URL: ", url , " Service: ", service, " Name: ", name)
			a.setDownloadStatus("Adding Playlist", false, 1, 1)
		}
	default:
		log.Println("Error: Unsupported service in AddEntry:", service)	
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
	return nil
}
