package fsHandler

import (
	"log"
	"os"
	"strings"

	"golang.org/x/text/unicode/norm"
)

// This function should scan the library and populate the FileSystem struct
func ScanLibrary(path string, fs *FileSystem) error {
	
	/*if fs == nil {
        fs = &FileSystem{}
    }*/
	// In Go, when you pass a struct or pointer to a function, 
	// you are passing a copy of the pointer. 
	// If you reassign the pointer inside the function 
	// (e.g., fs = &FileSystem{}), it only updates the local copy 
	// of the pointer, not the original pointer in the caller.

	// Clear the previous directories
	fs.Directories = nil

	directories, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	var folders []os.DirEntry
	for _, entry := range directories {
		if entry.IsDir() {
			folders = append(folders, entry)
		}
	}
	directories = folders
	log.Println("Scanning library: ", len(directories), " folders found")

	for _, folder := range directories {
		//log.Println("-- Scanning Dir: ", folder.Name())
		fs.Directories = append(fs.Directories, Playlist{
			PlaylistName: folder.Name(),
		})
		files, _ := os.ReadDir(path + "/" + folder.Name())
		for _, file := range files {
			if !file.IsDir() {
				//log.Println("---- File Found: ", file.Name())
				fs.Directories[len(fs.Directories)-1].File = append(fs.Directories[len(fs.Directories)-1].File, Song{
					Name:         file.Name(),
				})
			}
		}
	}

	// FUNCTION TO VERIFY IF THE FILES ARE IN THE FILESYSTEM DATA STRUCTURE
	log.Println("---- Filesystem structure: ----")
	for _, playlist := range fs.Directories {
		log.Println("Playlist:", playlist.PlaylistName)
		for _, song := range playlist.File {
			log.Println(" | Song:", song.Name)
		}
	}

	return nil
}

func GetSongsInFolder(fs *FileSystem, folderName string) ([]Song, error) {
	var songs []Song

	normalizedFolderName := norm.NFC.String(strings.TrimSpace(folderName))

	for _, playlist := range fs.Directories {
		normalizedPlaylistName := norm.NFC.String(strings.TrimSpace(playlist.PlaylistName))

		if normalizedPlaylistName == normalizedFolderName {
			log.Printf("Match found for '%s'. Found %d songs.\n", normalizedFolderName, len(playlist.File))
			songs = playlist.File
			break
		}
	}

	if len(songs) == 0 {
		log.Printf("Warning: No match found in FileSystem struct for folder name '%s'.\n", normalizedFolderName)
	}

	return songs, nil
}