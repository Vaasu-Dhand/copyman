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
		Width:  700, // Increased for keyboard feel
		Height: 450, // Better proportions for keyboard layout
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.OnStartup,
		OnShutdown:       app.OnShutdown,
		Frameless:        true, // Keep frameless but add custom titlebar
		StartHidden:      true,
		WindowStartState: options.Minimised, // Start hidden
		AlwaysOnTop:      false,             // We control this programmatically
		DisableResize:    false,             // Allow resize for better UX
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,  // Transparent titlebar
				HideTitle:                  true,  // Hide the title text
				HideTitleBar:               false, // Show titlebar area
				FullSizeContent:            true,  // Content extends under titlebar
				UseToolbar:                 true,  // Use toolbar style
				HideToolbarSeparator:       true,  // Hide separator line
			},
			Appearance:           mac.NSAppearanceNameAqua,
			WebviewIsTransparent: true,  // Keep transparent
			WindowIsTranslucent:  false, // Window background opaque
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
