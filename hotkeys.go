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
	log.Println("Starting hotkey listener...")

	// Register Cmd+Shift+C to show overlay
	hook.Register(hook.KeyDown, []string{"cmd", "shift", "c"}, func(e hook.Event) {
		log.Println("Hotkey triggered: Cmd+Shift+C")
		h.app.ShowOverlay()
	})

	// Register Escape to hide overlay when window is focused
	hook.Register(hook.KeyDown, []string{"escape"}, func(e hook.Event) {
		log.Println("Escape key pressed")
		h.app.HideOverlay()
	})

	// Register number keys 1-9 with Cmd+Shift for quick copy
	for i := 1; i <= 9; i++ {
		keyStr := strconv.Itoa(i)
		// Capture the current value of keyStr for the closure
		capturedKey := keyStr
		hook.Register(hook.KeyDown, []string{"cmd", "shift", capturedKey}, func(e hook.Event) {
			settings := h.app.GetSettings()
			if text, exists := settings.KeyBindings[capturedKey]; exists && text != "" {
				log.Printf("Quick copy triggered for key %s: %s", capturedKey, text)
				h.app.CopyToClipboard(text)
			}
		})
	}

	// Start the hook event loop
	s := hook.Start()
	<-hook.Process(s)
}

// StopListening stops the hotkey listener
func (h *HotkeyManager) StopListening() {
	if h.running {
		hook.End()
		h.running = false
		log.Println("Hotkey listener stopped")
	}
}
