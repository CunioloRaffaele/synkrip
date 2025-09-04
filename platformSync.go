package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"synkrip/api"
    _"synkrip/api/spotify" // import must be added to allow spotify to register with init()
	"synkrip/api/youtube"
	"synkrip/fsHandler"
	"strings"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

// normalizeString pulisce una stringa per il confronto.
// Rimuove l'estensione, converte in minuscolo e rimuove i caratteri non alfanumerici.
func normalizeString(s string) string {
    // 1. Rimuovi l'estensione del file (es. ".m4a", ".mp3")
    s = strings.TrimSuffix(s, filepath.Ext(s))

    // 2. Converti tutto in minuscolo
    s = strings.ToLower(s)

    // 3. Rimuovi tutti i caratteri che non sono lettere o numeri
    // Questo elimina spazi, trattini, parentesi, ecc.
    reg := regexp.MustCompile("[^a-z0-9]+")
    s = reg.ReplaceAllString(s, "")

    return s
}

// compareSongNames confronta in modo flessibile un nome di file dal filesystem
// con un nome di canzone e artista dal database.
// Restituisce true se c'è una corrispondenza probabile.
func compareSongNames(fsName, dbSongName, dbArtistName string) bool {
    normFsName := normalizeString(fsName)
    normDbSong := normalizeString(dbSongName)
    normDbArtist := normalizeString(dbArtistName)

    // Caso 1: Il nome del file normalizzato contiene sia la canzone che l'artista.
    // Questo è il caso più affidabile (es. "artistmysong.m4a")
    if strings.Contains(normFsName, normDbSong) && strings.Contains(normFsName, normDbArtist) {
        return true
    }

    // Caso 2: Il nome del file normalizzato corrisponde solo alla canzone.
    // Meno affidabile, ma utile se l'utente ha omesso l'artista.
    if strings.Contains(normFsName, normDbSong) {
        return true
    }

    return false
}

func (a *App) updatePlaylistDb() error {
    // Get all playlists from the database
    playlists, err := a.CurrentDB.GetPlaylist()
    if err != nil {
        log.Println("Error getting playlists:", err)
        return err
    }

    // Iterate through each playlist
    for index, pl := range playlists {

        if (pl.SERVICE == "unknownService") {
            // this happens when user adds a new playlist folder which is not associated to a service
            // SKIP
            continue
        }

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

        // Create a map of song IDs in the service playlist
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
                response, _ := rt.MessageDialog(a.ctx, rt.MessageDialogOptions{
                    Title:   "Song Removed from Service",
                    Message: fmt.Sprintf("The song '%s' by '%s' was removed from the playlist '%s' on the service. Do you want to delete the local file?", song.SONG_NAME, song.SONG_ARTIST_NAME, pl.DIR_ID),
                    Type:    "question",
                    Buttons: []string{"Delete File", "Keep File"},
                    DefaultButton: "Keep File",
                })
                if response == "Delete File" {
                    err = fsHandler.RemoveSongFile(a.LibPath, pl.DIR_ID, song.SONG_NAME + " - " + song.SONG_ARTIST_NAME + ".m4a")
                }
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

        // Verify if the folder contains music that was imported externally and if db need to be updated with "downloaded = 1"
        songsDbArr, err = a.CurrentDB.GetSongs(pl.DIR_ID)
        if err != nil {
            log.Println("Error getting songs for playlist:", pl.DIR_ID, err)
            return err
        }
        songsFsArr, err := fsHandler.GetSongsInFolder(a.CurrentFileSystem, pl.DIR_ID)
        if err != nil {
            log.Println("Error getting songs from filesystem for playlist:", pl.DIR_ID, err)
            return err
        }
        // now compare songsDbArr and songsFsArr to find discrepancies
        // if there is some song in fsArr that are also in dbArr (as not downloaded) they need to be marked as Downloaded
        for _, songFs := range songsFsArr {
            for _, songDb := range songsDbArr {
                if compareSongNames(songFs.Name, songDb.SONG_NAME, songDb.SONG_ARTIST_NAME) {
                    log.Printf("Marking song '%s' as downloaded in playlist '%s'\n", songFs.Name, pl.DIR_ID)
                    a.CurrentDB.MarkSongAsDownloaded(songDb.DIR_ID, songDb.SONG_YT_ID)
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
