package main

import (
	"context"
	"fmt"
	"log"
	"synkrip/api"
    _"synkrip/api/spotify" // import must be added to allow spotify to register with init()
	"synkrip/api/youtube"
	"synkrip/fsHandler"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) updatePlaylistDb() error {
    // Get all playlists from the database
    playlists, err := a.CurrentDB.GetPlaylist()
    if err != nil {
        log.Println("Error getting playlists:", err)
        return err
    }

    // Iterate through each playlist
    for index, pl := range playlists {

        // display download status in the frontend
        a.setDownloadStatus("Db Update", true, index, len(playlists))

        // Get all songs in the database for the current playlist
        songsDbArr, err := a.CurrentDB.GetSongs(pl.DIR_ID)
        if err != nil {
            log.Println("Error getting songs for playlist:", pl.DIR_ID, err)
            return err
        }

        // Create a map of song IDs in the database for quick lookup
        songIdsInDB := make(map[string]bool)
        for _, song := range songsDbArr {
            songIdsInDB[song.SONG_PLATFORM_ID] = true
        }

        // Handle synchronization based on the service
        musicService, err := api.GetMusicService(pl.SERVICE)
        if err != nil {
            log.Println("Error getting music service:", err)
            a.setDownloadStatus("", false, 0, 0)
            return err
        }

        playlistData, err := musicService.FetchPlaylist(pl.URL)
        if err != nil {
            log.Println("Error fetching playlist data:", err)
            a.setDownloadStatus("", false, 0, 0)
            return err
        }

        // update database image
        err = a.CurrentDB.UpdateImage(pl.DIR_ID, playlistData.Image)
        if err != nil {
            log.Println("Error updating playlist image:", err)
            a.setDownloadStatus("", false, 0, 0)
            return err
        }

        // Create a map of song IDs in the Spotify playlist
        songIdsInService := make(map[string]bool)
        for _, item := range playlistData.Songs {
            songIdsInService[item.ID] = true
        }

        // Remove songs from the database that are no longer in the service playlist
        for _, song := range songsDbArr {
            if !songIdsInService[song.SONG_PLATFORM_ID] {
                log.Printf("Removing song '%s' from playlist '%s' as it is no longer in service\n", song.SONG_NAME, pl.DIR_ID)
                var err error
                err = a.CurrentDB.RemoveSongFromPlaylist(pl.DIR_ID, song.SONG_PLATFORM_ID)
                if err != nil {
                    log.Printf("Error removing song '%s' from db '%s': %v\n", song.SONG_NAME, pl.DIR_ID, err)
                    a.setDownloadStatus("", false, 0, 0)
                    return err
                }
                err = fsHandler.RemoveSongFile(a.LibPath, pl.DIR_ID, song.SONG_NAME + " - " + song.SONG_ARTIST_NAME + ".m4a")
                if err != nil {
                    log.Printf("Error removing song '%s' from filesystem '%s': %v\n", song.SONG_NAME, pl.DIR_ID, err)
                    a.setDownloadStatus("", false, 0, 0)
                    return err
                }
                
            }
        }

        // Add new songs from the service playlist to the database
        for _, item := range playlistData.Songs {
            if !songIdsInDB[item.ID] {
                log.Printf("Adding new song '%s' to playlist '%s'\n", item.Name, pl.DIR_ID)
                ytId, _ := youtube.GetYTid(item.Name + " " + item.Artists[0])
                err := a.CurrentDB.AddSongInPlaylist(
                    pl.DIR_ID,
                    item.Name,
                    item.Album,
                    item.Artists[0],
                    ytId,
                    item.ID,
                    false, // IS_DOWNLOADED
                )
                if err != nil {
                    log.Printf("Error adding song '%s' to playlist '%s': %v\n", item.Name, pl.DIR_ID, err)
                    a.setDownloadStatus("", false, 0, 0)
                    return err
                }
            }
        }
    }

    log.Println("Playlist synchronization completed.")
    // Emit event to hide download status in frontend
    a.setDownloadStatus("", false, 0, 0)
    return nil
}

func (a *App) downloadContent(ctx context.Context, playlistName string) {
    // Get all songs in the playlist
    songs, err := a.CurrentDB.GetSongs(playlistName)
    if err != nil {
        log.Println("Error getting songs for playlist:", playlistName, err)
        return
    }

    a.setDownloadStatus("Downloading...", true, 0, len(songs))

    // Iterate through each song and download it
    for index, song := range songs {
        // Check if the context is canceled
        select {
        case <-ctx.Done():
            log.Println("Download canceled for playlist:", playlistName)
            a.setDownloadStatus("", false, 0, 0)
            return
        default:
            // Continue if not canceled
        }
        a.setDownloadStatus("Downloading...", true, index, len(songs))
        if song.IS_DOWNLOADED == 0 {
            log.Printf("Downloading song '%s' from playlist '%s'\n", song.SONG_NAME, playlistName)
            err := a.DownloadVideo(song.SONG_YT_ID, a.LibPath+"/"+playlistName, song.SONG_NAME + " - " + song.SONG_ARTIST_NAME)
            if (err != nil) {
                log.Printf("Download aborted")
                a.setDownloadStatus("", false, 0, 0)
                rt.MessageDialog(a.ctx, rt.MessageDialogOptions{
                    Title:   "Download Error",
                    Message: fmt.Sprintf("Failed to download song '%s': %v", song.SONG_NAME, err),
                    Type:    "error",
                })
                return
            } else {
                a.CurrentDB.MarkSongAsDownloaded(song.DIR_ID, song.SONG_YT_ID)
            }        
        }
    }
    log.Println("Download completed for playlist:", playlistName)
    a.setDownloadStatus("", false, 0, 0)

}
