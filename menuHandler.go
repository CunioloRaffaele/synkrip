package main

import (
	"log"
	"runtime"
	"synkrip/fsHandler"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

// Create the File menu
func createFileMenu(app *App) *menu.Menu {
    return menu.NewMenuFromItems(
        menu.Text("Check for app updates", nil, func(_ *menu.CallbackData) {
            checkForUpdate(app)
        }),
        menu.Text("View open source libs", nil, func(_ *menu.CallbackData) {
            rt.EventsEmit(app.ctx, "osLib")
        }),
        menu.Text("Settings", nil, func(_ *menu.CallbackData) {
            rt.EventsEmit(app.ctx, "settings")
        }),
        menu.Separator(),
        menu.Text("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
            rt.Quit(app.ctx)
        }),
    )
}

// Create the Library menu
func createLibraryMenu(app *App) *menu.Menu {
    return menu.NewMenuFromItems(
        menu.Text("Open Library", keys.CmdOrCtrl("l"), func(_ *menu.CallbackData) {
            err := fsHandler.OpenLibrary(app.ctx, &app.LibPath, &app.CurrentDB, app.CurrentFileSystem, "")
			if err != nil {
				log.Println("Failed to open library:", err.Error())
			} else {
                app.Settings.AddLibrary(app.LibPath)
            }
        }),
        menu.Text("Create New Library", keys.CmdOrCtrl("n"), func(_ *menu.CallbackData) {
            err := fsHandler.CreateLibrary(app.ctx, &app.LibPath, &app.CurrentDB)
            if err != nil {
				log.Println("Failed to create library:", err.Error())
			} else {
                app.Settings.AddLibrary(app.LibPath)
            }
        }),
    )
}

// Create the main app menu
func createAppMenu(app *App) *menu.Menu {
    appMenu := menu.NewMenu()

    if runtime.GOOS == "darwin" {
        appMenu.Append(menu.AppMenu()) // macOS-specific menu
		appMenu.Append(menu.EditMenu())  // On macOS platform, EditMenu should be appended to enable Cmd+C, Cmd+V, Cmd+Z... shortcuts
    }

    // Append the File menu as a submenu
    appMenu.Append(menu.SubMenu("File", createFileMenu(app)))

    // Append the Library menu as a submenu
    appMenu.Append(menu.SubMenu("Library", createLibraryMenu(app)))

    return appMenu
}