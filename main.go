package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:     "ADB Commander Blue Print",
		Width:     1000,
		Height:    620,
		MinWidth:  800,
		MinHeight: 600,
		// Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []any{
			app,
		},
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop: true,
		},
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId: "adb-commander-blueprint-unique-id",
		},
		// Windows: &windows.Options{
		// 	WebviewIsTransparent: true,
		// 	WindowIsTranslucent:  true,
		// 	BackdropType:         windows.Mica,
		// 	Theme:                windows.Dark,
		// },
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
