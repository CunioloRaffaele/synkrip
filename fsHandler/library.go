package fsHandler

import (
    "log"
    "strings"
    "synkrip/dbHandler"
	"context"

    rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func OpenLibrary(ctx context.Context, libPath *string, currentDB **dbHandler.Database, currentFileSystem *FileSystem, newLib string) error {
    var err error

    // If newLib is empty, open a directory dialog to select the library path
    if len(newLib) == 0 {
        *libPath, err = rt.OpenDirectoryDialog(ctx, rt.OpenDialogOptions{})
        if err != nil {
            log.Println("Error selecting directory:", err.Error())
            return err
        }
    } else {
        // If newLib is provided, use it as the library path
        *libPath = newLib
    }

    if len(*libPath) > 0 {
        folderName := (*libPath)[strings.LastIndex(*libPath, "/")+1:]
        log.Println("Selected folder:", folderName)
        rt.WindowSetTitle(ctx, "SynkRip - "+folderName)

        var errDb error
        *currentDB, errDb = dbHandler.Connect(*libPath)
        if errDb != nil {
            rt.MessageDialog(ctx, rt.MessageDialogOptions{
                Type:    rt.WarningDialog,
                Title:   "Error",
                Message: "Failed to open library:\n" + errDb.Error(),
            })
            return errDb
        }

        fsErr := ScanLibrary(*libPath, currentFileSystem)
        log.Println("Library scan complete. File system:", currentFileSystem)
        if fsErr != nil {
            rt.MessageDialog(ctx, rt.MessageDialogOptions{
                Type:    rt.WarningDialog,
                Title:   "Error",
                Message: "Failed to scan library:\n" + fsErr.Error(),
            })
            return fsErr
        }
        rt.EventsEmit(ctx, "LibSelected")
        
    }

    return nil
}

func CreateLibrary(ctx context.Context, libPath *string, currentDB **dbHandler.Database) error {
    var err error
    *libPath, err = rt.OpenDirectoryDialog(ctx, rt.OpenDialogOptions{})
    if err != nil {
        log.Println("Error selecting directory:", err.Error())
        return err
    }

    if len(*libPath) > 0 {
        folderName := (*libPath)[strings.LastIndex(*libPath, "/")+1:]
        log.Println("Selected folder:", folderName)
        rt.WindowSetTitle(ctx, "SynkRip - "+folderName)

        err = dbHandler.CreateDatabase(*libPath)
        if err != nil {
            rt.MessageDialog(ctx, rt.MessageDialogOptions{
                Type:    rt.WarningDialog,
                Title:   "Error",
                Message: "Failed to create library:\n" + err.Error(),
            })
            return err
        }

        var errDb error
        *currentDB, errDb = dbHandler.Connect(*libPath)
        if errDb != nil {
            log.Println("Error connecting to database:", errDb.Error())
            return errDb
        }

        (*currentDB).Setup()
        rt.EventsEmit(ctx, "LibSelected")
    }

    return nil
}