package main

import rt "github.com/wailsapp/wails/v2/pkg/runtime"

func (app *App)setDownloadStatus(title string, isVisible bool, currentItem int, totalItems int) {
	// Emit event to show download status in frontend
    rt.EventsEmit(app.ctx, "updateDownloadProgress", map[string]interface{}{
        "title":        title,
        "isVisible":   	isVisible,
        "currentItem": 	currentItem,
        "totalItems": 	totalItems,
    })
}