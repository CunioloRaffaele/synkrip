package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"synkrip/fsHandler"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
    ytdlpPath   string
    ffmpegPath  string
    ffprobePath string
)

// extractBinary extracts an embedded binary to the filesystem and makes it executable
func extractBinary(binaryName string, binaryData []byte) (string, error) {
    path, err := fsHandler.GetAppDataPath()
    if err != nil {
        log.Printf("Error getting app data path for %s: %v", binaryName, err)
        return "", err
    }
    
    binaryLocation := filepath.Join(path, binaryName)

    // Check if the file already exists
    if _, err := os.Stat(binaryLocation); err == nil {
        log.Printf("%s binary already exists, skipping extraction", binaryName)
    } else {
        log.Printf("Extracting %s binary to: %s", binaryName, binaryLocation)
        err = os.WriteFile(binaryLocation, binaryData, 0755) // Make it executable
        if err != nil {
            log.Printf("Error extracting %s binary: %v", binaryName, err)
            return "", err
        }
    }

    return binaryLocation, nil
}

// TODO : add multi-platform support for binaries
// externalFrameworksInit extracts all required binaries
func externalFrameworksInit() {
	var err error
    // Extract yt-dlp binary
    ytdlpPath, err = extractBinary("ytdlp_darwin", ytdlp_darwin)
    if err != nil {
        log.Printf("Failed to extract yt-dlp: %v", err)
    }
    
    // Extract ffmpeg binary
    ffmpegPath, err = extractBinary("ffmpeg_darwin", ffmpeg_darwin)
    if err != nil {
        log.Printf("Failed to extract ffmpeg: %v", err)
    }
    
    // Extract ffprobe binary
    ffprobePath, err = extractBinary("ffprobe_darwin", ffprobe_darwin)
    if err != nil {
        log.Printf("Failed to extract ffprobe: %v", err)
    }

    log.Printf("Binaries extracted: yt-dlp=%s, ffmpeg=%s, ffprobe=%s", 
        ytdlpPath, ffmpegPath, ffprobePath)
}


// DownloadVideo downloads a video using yt-dlp without blocking the main thread
func (a *App) DownloadVideo(url, outputDir string) {
    log.Printf("Starting download from URL: %s to directory: %s", url, outputDir)

        // Set up command with options for audio extraction
        cmd := exec.Command(
            ytdlpPath,
            "--extract-audio",
            "--audio-format", "m4a",
            "--audio-quality", "0", // 0 is best
            "--output", filepath.Join(outputDir, "%(title)s.%(ext)s"),
			"--embed-thumbnail",
			"--write-thumbnail",
			"--ffmpeg-location", ffmpegPath,
            /*"--progress-template", "download:%(progress.downloaded_bytes)s/%(progress.total_bytes)s",*/
            url,
        )

        cmd.Stdout = log.Writer()
        cmd.Stderr = log.Writer()

        // Start the command
        if err := cmd.Run(); err != nil {
            log.Printf("Error starting yt-dlp: %v", err)
            rt.EventsEmit(a.ctx, "updateDownloadProgress", map[string]interface{}{
                "isVisible": false,
            })
            return
        }

        // Wait for command to finish
        _ = cmd.Wait()
        
        // Command completed - hide download status
        rt.EventsEmit(a.ctx, "updateDownloadProgress", map[string]interface{}{
            "isVisible": false,
        })

        log.Println("Download completed successfully")
}
