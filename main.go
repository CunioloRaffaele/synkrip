package main

import (
	"embed"
	"log"
	"runtime"
	"strings"

	"synkrip/dbHandler"
	"synkrip/fsHandler"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed wails.json
var wailsJSON []byte

// TODO : add multi-platform support for binaries
//go:embed platformBinary/yt-dlp_darwin
var ytdlp_darwin []byte
//go:embed platformBinary/ffmpeg_darwin
var ffmpeg_darwin []byte
//go:embed platformBinary/ffprobe_darwin
var ffprobe_darwin []byte


func main() {

	// Create an instance of the App struct
    app := NewApp()

	AppMenu := menu.NewMenu()
    if runtime.GOOS == "darwin" {
        AppMenu.Append(menu.AppMenu()) // On macOS platform, this must be done right after `NewMenu()`
    }
    FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("Check for app updates", nil, func(_ *menu.CallbackData) {
		checkForUpdate(app)
    })
	FileMenu.AddText("View open source libs", nil, func(_ *menu.CallbackData) {
		rt.EventsEmit(app.ctx, "osLib")
    })
    FileMenu.AddSeparator()
    FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
        rt.Quit(app.ctx)
    })
	LibraryMenu := AppMenu.AddSubmenu("Library")
	LibraryMenu.AddText("Open Library", keys.CmdOrCtrl("l"), func(_ *menu.CallbackData) {
		var err error
		app.LibPath, err = rt.OpenDirectoryDialog(app.ctx, rt.OpenDialogOptions{})
		if err != nil{
			println("Error selecting directory:", err.Error())
		} else if (len(app.LibPath) > 0){
			folderName := app.LibPath[strings.LastIndex(app.LibPath, "/")+1:]
			log.Println("Selected folder: " + folderName)
			rt.WindowSetTitle(app.ctx,"SynkRip - " + folderName)
			var errDb error
			app.CurrentDB, errDb = dbHandler.Connect(app.LibPath)
			if errDb != nil {
				rt.MessageDialog(app.ctx, rt.MessageDialogOptions{
				Type:          rt.WarningDialog,
				Title:         "Error",
				Message:       "Failed to open library:\n" + errDb.Error(),
				})
			} else {
				fsErr := fsHandler.ScanLibrary(app.LibPath, app.CurrentFileSystem)
				log.Println("Library scan complete. File system:", app.CurrentFileSystem)
				if fsErr != nil {
					rt.MessageDialog(app.ctx, rt.MessageDialogOptions{
					Type:          rt.WarningDialog,
					Title:         "Error",
					Message:       "Failed to scan library:\n" + fsErr.Error(),
					})
				} else {
					rt.EventsEmit(app.ctx, "LibSelected")
					app.updatePlaylistDb()
				}
			}
		}
	})
	LibraryMenu.AddText("Create New Library", keys.CmdOrCtrl("a"), func(_ *menu.CallbackData) {
		var err error
		app.LibPath, err = rt.OpenDirectoryDialog(app.ctx, rt.OpenDialogOptions{})
		if err != nil{
			println("Error selecting directory:", err.Error())
		} else if (len(app.LibPath) > 0){
			folderName := app.LibPath[strings.LastIndex(app.LibPath, "/")+1:]
			log.Println("Selected folder: " + folderName)
			rt.WindowSetTitle(app.ctx,"SynkRip - " + folderName)
			err := dbHandler.CreateDatabase(app.LibPath)
			if err != nil {
				rt.MessageDialog(app.ctx, rt.MessageDialogOptions{
				Type:          rt.WarningDialog,
				Title:         "Error",
				Message:       "Failed to create library:\n" + err.Error(),
				})
			}
			var errDb error
			app.CurrentDB, errDb = dbHandler.Connect(app.LibPath)
			if errDb != nil {
				println("Error connecting to database:", errDb.Error())
			} else {
				app.CurrentDB.Setup()
				rt.EventsEmit(app.ctx, "LibSelected")
			}
		}
	})
	/*LibraryMenu.AddText("Scan", keys.CmdOrCtrl("s"), func(_ *menu.CallbackData) {
		if app.LibPath != "" {
			fsErr := fsHandler.ScanLibrary(app.LibPath, app.CurrentFileSystem)
			if fsErr != nil {
				rt.MessageDialog(app.ctx, rt.MessageDialogOptions{
				Type:          rt.WarningDialog,
				Title:         "Error",
				Message:       "Failed to scan library:\n" + fsErr.Error(),
				})
			} else {
				rt.EventsEmit(app.ctx, "LibSelected")
			}
		} else {
			rt.MessageDialog(app.ctx, rt.MessageDialogOptions{
			Type:          rt.WarningDialog,
			Title:         "Error",
			Message:       "No library selected.",
			})
		}
	
	})*/


    if runtime.GOOS == "darwin" {
    AppMenu.Append(menu.EditMenu())  // On macOS platform, EditMenu should be appended to enable Cmd+C, Cmd+V, Cmd+Z... shortcuts
    }

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "synkrip APP",
		Menu: AppMenu, 
		Width:  1024,
		Height: 768,
		/*MinWidth: 1024,
        MinHeight: 768,*/
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown: 	  app.shutdown,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
            TitleBar: &mac.TitleBar{
                TitlebarAppearsTransparent: false,
                HideTitle:                  false,
                HideTitleBar:               false,
                FullSizeContent:            false,
                UseToolbar:                 false,
                HideToolbarSeparator:       true,
            },
            Appearance:           mac.NSAppearanceNameDarkAqua,
            WebviewIsTransparent: true,
            WindowIsTranslucent:  true,
            About: &mac.AboutInfo{
                Title:   "SynkRip",
                Message: "Made with ❤️ by Cuniolo Raffaele.\n\nPlease use it wisely and support the artists who create the songs you download.\n\n" + "Version: " + getCurrentVersion() + "\n" + "Compiled: " + getCompilationDate(),
            },
			DisableZoom: true,
        },
		
	})

	if err != nil {
		println("Error:", err.Error())
	}

}