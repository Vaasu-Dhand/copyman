package main

import (
	"log"
	"strconv"

	hook "github.com/robotn/gohook"
)

// HotkeyManager handles global hotkeys
type HotkeyManager struct {
	app     *App
	running bool
}

// NewHotkeyManager creates a new hotkey manager
func NewHotkeyManager(app *App) *HotkeyManager {
	return &HotkeyManager{
		app:     app,
		running: false,
	}
}

// StartListening starts listening for global hotkeys
func (h *HotkeyManager) StartListening() {
	if h.running {
		return
	}

	h.running = true
	log.Println("Copyman: Starting hotkey listener...")

	// Register Ctrl+Shift+Space to toggle overlay
	hook.Register(hook.KeyDown, []string{"ctrl", "shift", "space"}, func(e hook.Event) {
		log.Println("Copyman: Toggle hotkey triggered (Ctrl+Shift+Space)")
		h.app.ToggleOverlay()
	})

	// DON'T register global Escape - let the frontend handle it when window has focus

	// Register number keys 1-9 with Cmd+Shift for quick copy with flash
	for i := 1; i <= 9; i++ {
		keyStr := strconv.Itoa(i)
		// Capture the current value of keyStr for the closure
		capturedKey := keyStr
		hook.Register(hook.KeyDown, []string{"cmd", "shift", capturedKey}, func(e hook.Event) {
			if h.app.IsVisible() { // Only work when overlay is visible
				settings := h.app.GetSettings()
				if text, exists := settings.KeyBindings[capturedKey]; exists && text != "" {
					log.Printf("Copyman: Quick copy key %s triggered", capturedKey)
					h.app.CopyAndFlash(text, capturedKey)
				}
			}
		})
	}

	log.Println("Copyman: Hotkeys registered - Ctrl+Shift+Space to toggle, Cmd+Shift+1-9 for quick copy")
	log.Println("Copyman: Escape key handled by UI when window has focus")

	// Start the hook event loop
	s := hook.Start()
	<-hook.Process(s)
}

// StopListening stops the hotkey listener
func (h *HotkeyManager) StopListening() {
	if h.running {
		hook.End()
		h.running = false
		log.Println("Copyman: Hotkey listener stopped")
	}
}
