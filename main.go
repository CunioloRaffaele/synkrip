package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
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

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "synkrip APP",
		Menu: createAppMenu(app),
		Width:  1024,
		Height: 700,
		MinWidth: 1024,
        MinHeight: 700,
		DisableResize: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnDomReady:         app.domready,
		OnShutdown: 	  app.shutdown,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
            TitleBar: &mac.TitleBar{
                TitlebarAppearsTransparent: false,
                HideTitle:                  false,
                HideTitleBar:               false,
                FullSizeContent:            true,
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