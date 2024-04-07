package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"kuddle/build"
	"kuddle/service/pkg/kuddle"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := kuddle.NewApp()
	build.Icon()

	// Create application with options
	err := wails.Run(&options.App{
		Title: "Kuddle",
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.Startup,
		Bind: []interface{}{
			app,
		},
		Frameless:        true,
		Height:           720,
		Width:            720,
		MaxHeight:        720,
		MaxWidth:         720,
		WindowStartState: options.Normal,
		Fullscreen:       false,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               "kuddle",
			OnSecondInstanceLaunch: nil,
		},
		Linux: &linux.Options{
			Icon: icon,
		},
		DisableResize: true,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
