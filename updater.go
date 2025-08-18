package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/tidwall/gjson"
	r "github.com/wailsapp/wails/v2/pkg/runtime"
)

const repoURL = "https://api.github.com/repos/CunioloRaffaele/synkrip/releases/latest"

type Release struct {
	TagName string `json:"tag_name"`
}

func getCurrentVersion() (string) {
	version:= gjson.Get(string(wailsJSON) , "info.productVersion")
	log.Println("Currently running: " + version.String())
	log.Println("OS: " + runtime.GOOS + " " + runtime.GOARCH)
	return version.String()
}

func getCompilationDate() (string) {
	version:= gjson.Get(string(wailsJSON) , "info.compilationDate")
	log.Println("Compiled: " + version.String())
	return version.String()
}

func checkForUpdate(app *App) {

	version:= getCurrentVersion()

	if !app.Settings.UpdateCheck {
		log.Println("Update check is disabled in settings, skipping update check.")
		return
	}

	fmt.Println("Checking for updates...")
	resp, err := http.Get(repoURL)
	if err != nil {
		log.Printf("failed to get latest release: %v", err)
		r.MessageDialog(app.ctx, r.MessageDialogOptions{
			Type:          r.ErrorDialog,
			Title:         "Error",
			Message:       "Cannot connect to update service. Verify connection...",
			Buttons:       []string{"Ok"},
			DefaultButton: "Ok",
			})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		return
	}
	latestVersion:= gjson.Get(string(body) , "tag_name")
	log.Println("Latest Version: " + (latestVersion).String())
	latestVersionUrl:= gjson.Get(string(body) , "html_url")
	log.Println("Latest Version Download URL: " + (latestVersionUrl).String())

	if (strings.Compare(version, latestVersion.String()) == 0) || (len(latestVersion.String()) == 0) {
		log.Println("You are running the latest version.")
	} else {
		log.Println("There is a new version available.")
		result, err := r.MessageDialog(app.ctx, r.MessageDialogOptions{
        Type:          r.WarningDialog,
        Title:         "Update Available",
        Message:       "Do you want to download the latest version?" + "\n\n" + "The latest Version is: " + latestVersion.String(),
        Buttons:       []string{"Yes", "No"},
		DefaultButton: "Yes",
    	})
		if err != nil {
			log.Println(err)
		}
		if result == "Yes" {
			// Open the browser to the latest release
			var err error
			switch os := runtime.GOOS; os {
			case "windows":
				err = exec.Command("rundll32", "url.dll,FileProtocolHandler", latestVersionUrl.String()).Start()
			case "darwin":
				err = exec.Command("open", latestVersionUrl.String()).Start()
			}
			if err != nil {
				log.Println(err)
			}
		} else {
			r.MessageDialog(app.ctx, r.MessageDialogOptions{
			Type:          r.ErrorDialog,
			Title:         "Warning",
			Message:       "This version may not be supported anymore. \nIf you encounter bugs please update.",
			Buttons:       []string{"Ok"},
			DefaultButton: "Ok",
			})
		}
	}

}
