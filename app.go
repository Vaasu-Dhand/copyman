package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
)

// App struct
type App struct {
	ctx            context.Context
	settings       *Settings
	hotkeyManager  *HotkeyManager
	isVisible      bool
	lastToggleTime int64 // Timestamp to prevent rapid toggling
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
	a.isVisible = false
	a.lastToggleTime = 0

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
	// Stop hotkey listener
	if a.hotkeyManager != nil {
		a.hotkeyManager.StopListening()
	}

	// Save settings before shutdown
	a.saveSettings()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, welcome to Copyman!", name)
}

// GetSettings returns the current settings
func (a *App) GetSettings() Settings {
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
	return *a.settings
}

// SaveSettings saves the settings to file
func (a *App) SaveSettings(settings Settings) error {
	a.settings = &settings
	return a.saveSettings()
}

// CopyToClipboard copies text to the system clipboard
func (a *App) CopyToClipboard(text string) error {
	if text == "" {
		return fmt.Errorf("empty text")
	}

	clipboard.Write(clipboard.FmtText, []byte(text))
	log.Printf("Copyman: Text copied to clipboard (length: %d chars)", len(text))

	return nil
}

// CopyAndFlash copies text and triggers visual feedback
func (a *App) CopyAndFlash(text string, keyNumber string) error {
	if text == "" {
		return fmt.Errorf("empty text")
	}

	clipboard.Write(clipboard.FmtText, []byte(text))
	log.Printf("Copyman: Key %s copied to clipboard (length: %d chars)", keyNumber, len(text))

	// Emit event to frontend for visual feedback
	runtime.EventsEmit(a.ctx, "key-flashed", keyNumber)

	return nil
}

// ShowOverlay shows the main overlay window
func (a *App) ShowOverlay() {
	if !a.isVisible {
		runtime.WindowShow(a.ctx)
		runtime.WindowSetAlwaysOnTop(a.ctx, true)
		runtime.WindowCenter(a.ctx)
		a.isVisible = true
		log.Println("Copyman: Overlay window shown")
	}
}

// HideOverlay hides the main overlay window
func (a *App) HideOverlay() {
	if a.isVisible {
		runtime.WindowHide(a.ctx)
		a.isVisible = false
		log.Println("Copyman: Overlay window hidden")
	}
}

// ToggleOverlay toggles the overlay visibility with debouncing
func (a *App) ToggleOverlay() {
	currentTime := time.Now().UnixMilli()

	// Prevent rapid toggling (debounce 300ms)
	if currentTime-a.lastToggleTime < 300 {
		log.Println("Copyman: Toggle ignored - too rapid")
		return
	}

	a.lastToggleTime = currentTime

	if a.isVisible {
		a.HideOverlay()
	} else {
		a.ShowOverlay()
	}
}

// IsVisible returns whether the window is currently visible
func (a *App) IsVisible() bool {
	return a.isVisible
}

// getSettingsPath returns the path to the settings file
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
		defaultSettings := a.GetSettings()
		a.settings = &defaultSettings
		return
	}

	var settings Settings
	err = json.Unmarshal(data, &settings)
	if err != nil {
		log.Printf("Failed to unmarshal settings: %v", err)
		defaultSettings := a.GetSettings()
		a.settings = &defaultSettings
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
	a.hotkeyManager = NewHotkeyManager(a)

	// Start listening in a separate goroutine
	go func() {
		a.hotkeyManager.StartListening()
	}()

	log.Println("Global hotkeys registered")
}

func (a *App) OnSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
	// Show the overlay if someone tries to launch again
	a.ShowOverlay()
}
