package fsHandler

import (
	"os"
	"path/filepath"
)

func RemoveSongFile(libpath string, playlistID string, songID string) error {
	dir := filepath.Join(libpath, playlistID, songID)
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
	return nil
}