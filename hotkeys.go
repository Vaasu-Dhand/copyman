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

	// Register number keys 1-9 with Ctrl+Shift for quick copy with flash
	for i := 1; i <= 9; i++ {
		keyStr := strconv.Itoa(i)
		// Capture the current value of keyStr for the closure
		capturedKey := keyStr
		hook.Register(hook.KeyDown, []string{"ctrl", "shift", capturedKey}, func(e hook.Event) {
			log.Printf("Copyman: Hotkey Ctrl+Shift+%s detected", capturedKey)

			// Always handle the event to prevent system beep
			settings := h.app.GetSettings()
			if text, exists := settings.KeyBindings[capturedKey]; exists && text != "" {
				log.Printf("Copyman: Quick copy key %s triggered with text", capturedKey)
				h.app.CopyAndFlash(text, capturedKey)

				// Hide the window after copying if it's visible
				if h.app.IsVisible() {
					log.Printf("Copyman: Hiding window after copying key %s", capturedKey)
					h.app.HideOverlay()
				}
			} else {
				log.Printf("Copyman: No text bound to key %s, but event consumed to prevent beep", capturedKey)

				// Still hide the window if visible and key was pressed
				if h.app.IsVisible() {
					log.Printf("Copyman: Hiding window after empty key %s press", capturedKey)
					h.app.HideOverlay()
				}
			}
		})
	}

	log.Println("Copyman: Hotkeys registered - Ctrl+Shift+Space to toggle, Ctrl+Shift+1-9 for quick copy")
	log.Println("Copyman: Window will auto-hide after copying")

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
