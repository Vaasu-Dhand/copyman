package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
)

// App struct
type App struct {
	ctx      context.Context
	settings *Settings
}

// Settings represents the application settings
type Settings struct {
	Theme       string            `json:"theme"`
	KeyBindings map[string]string `json:"keyBindings"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// OnStartup is called when the app starts up
func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx

	// Initialize clipboard
	err := clipboard.Init()
	if err != nil {
		log.Printf("Failed to initialize clipboard: %v", err)
	}

	// Load settings
	a.loadSettings()

	// Register global hotkeys
	a.registerHotkeys()
}

// OnShutdown is called when the app is about to quit
func (a *App) OnShutdown(ctx context.Context) {
	// Save settings before shutdown
	a.saveSettings()
}

// GetSettings returns the current settings
func (a *App) GetSettings() *Settings {
	if a.settings == nil {
		a.settings = &Settings{
			Theme: "light",
			KeyBindings: map[string]string{
				"1": "",
				"2": "",
				"3": "",
				"4": "",
				"5": "",
				"6": "",
				"7": "",
				"8": "",
				"9": "",
			},
		}
	}
	return a.settings
}

// SaveSettings saves the settings to file
func (a *App) SaveSettings(settings *Settings) error {
	a.settings = settings
	return a.saveSettings()
}

// CopyToClipboard copies text to the system clipboard
func (a *App) CopyToClipboard(text string) error {
	if text == "" {
		return fmt.Errorf("empty text")
	}

	clipboard.Write(clipboard.FmtText, []byte(text))

	// Hide the window after copying
	runtime.WindowHide(a.ctx)

	return nil
}

// ShowOverlay shows the main overlay window
func (a *App) ShowOverlay() {
	runtime.WindowShow(a.ctx)
	runtime.WindowSetAlwaysOnTop(a.ctx, true)
}

// HideOverlay hides the main overlay window
func (a *App) HideOverlay() {
	runtime.WindowHide(a.ctx)
}

// ToggleOverlay toggles the overlay visibility
func (a *App) ToggleOverlay() {
	// This will be called by the global hotkey
	runtime.WindowShow(a.ctx)
	runtime.WindowSetAlwaysOnTop(a.ctx, true)
}

// GetSettingsPath returns the path to the settings file
func (a *App) getSettingsPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Failed to get home directory: %v", err)
		return "copyman_settings.json"
	}

	configDir := filepath.Join(homeDir, ".config", "copyman")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		log.Printf("Failed to create config directory: %v", err)
		return "copyman_settings.json"
	}

	return filepath.Join(configDir, "settings.json")
}

// loadSettings loads settings from file
func (a *App) loadSettings() {
	settingsPath := a.getSettingsPath()

	data, err := os.ReadFile(settingsPath)
	if err != nil {
		// File doesn't exist, use defaults
		a.settings = a.GetSettings()
		return
	}

	var settings Settings
	err = json.Unmarshal(data, &settings)
	if err != nil {
		log.Printf("Failed to unmarshal settings: %v", err)
		a.settings = a.GetSettings()
		return
	}

	a.settings = &settings
}

// saveSettings saves settings to file
func (a *App) saveSettings() error {
	if a.settings == nil {
		return fmt.Errorf("no settings to save")
	}

	settingsPath := a.getSettingsPath()

	data, err := json.MarshalIndent(a.settings, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal settings: %v", err)
	}

	err = os.WriteFile(settingsPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write settings file: %v", err)
	}

	return nil
}

// registerHotkeys registers global hotkeys
func (a *App) registerHotkeys() {
	// Note: Global hotkey registration would require a platform-specific library
	// For macOS, you might want to use a library like github.com/robotn/gohook
	// This is a placeholder for the hotkey registration logic
	log.Println("Hotkey registration would be implemented here")
}
