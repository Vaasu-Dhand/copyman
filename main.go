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

// main is the entry point for the application
func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Copyman",
		Width:  700,
		Height: 450,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.OnStartup,
		OnShutdown:       app.OnShutdown,
		Frameless:        false, // Changed to false to show title bar
		StartHidden:      true,
		WindowStartState: options.Minimised,
		AlwaysOnTop:      false,
		DisableResize:    true, // Set to true since you want fixed 700x450
		// Add these for better menu bar app behavior
		HideWindowOnClose: true,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               "com.copyman.copyman",
			OnSecondInstanceLaunch: app.OnSecondInstanceLaunch,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: false, // Set to false to show visible title bar
				HideTitle:                  false, // Show the title
				HideTitleBar:               false, // Don't hide the title bar
				FullSizeContent:            false, // Don't use full size content
				UseToolbar:                 false, // Don't use toolbar
				HideToolbarSeparator:       true,  // Keep this true
			},
			Appearance:           mac.NSAppearanceNameAqua,
			WebviewIsTransparent: false, // Set to false for proper title bar display
			WindowIsTranslucent:  false,
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
