# Copyman

A modern, minimalistic clipboard manager for macOS built with Wails, React, and TypeScript.

## Features

- 🎨 **Modern Design**: Clean, minimalistic interface with light/dark theme support
- ⌨️ **Global Hotkeys**: Quick access with `Cmd+Shift+C` and number key shortcuts
- 📋 **9 Quick Slots**: Bind frequently used text to number keys 1-9
- 🔄 **Persistent Settings**: Your configurations are saved between sessions
- 👁️ **Overlay Mode**: Transparent overlay that doesn't interrupt your workflow
- 🎯 **Lightweight**: Built with native performance using Wails

## Installation

### Prerequisites

- Go 1.21 or later
- Node.js 16 or later
- Wails v2 CLI tool

### Install Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### Setup Project

1. Clone the repository:
```bash
git clone <your-repo-url>
cd copyman
```

2. Install dependencies:
```bash
make install
```

3. Run in development mode:
```bash
make dev
```

## Building

### For Development
```bash
make dev
```

### For Production
```bash
make build
```

### For macOS Distribution (Universal Binary)
```bash
make build-darwin
make package-mac
```

This creates a universal binary supporting both Intel and Apple Silicon Macs.

## Usage

### Global Hotkeys

- **`Cmd+Shift+C`**: Open/show the Copyman overlay
- **`Escape`**: Close/hide the overlay
- **`1-9`**: Copy the text bound to that number key (overlay must be visible)
- **`Cmd+Shift+1-9`**: Quick copy without opening overlay (when implemented)

### Setting Up Text Shortcuts

1. Open Copyman (Cmd+Shift+C)
2. Click the "Settings" button
3. Enter text for each number key (1-9)
4. Settings are automatically saved
5. Use the number keys to instantly copy your predefined text

### Examples

Bind commonly used text like:
- Email addresses
- Phone numbers
- Frequently used passwords
- Code snippets
- Standard responses
- Social media handles

## Configuration

Settings are stored in `~/.config/copyman/settings.json`:

```json
{
  "theme": "light",
  "keyBindings": {
    "1": "your-email@example.com",
    "2": "+1 (555) 123-4567",
    "3": "Thank you for your message!",
    ...
  }
}
```

## Architecture

- **Backend**: Go with Wails v2 framework
- **Frontend**: React 18 + TypeScript + Vite
- **Styling**: CSS with CSS custom properties for theming
- **Build**: Native desktop app compilation
- **Hotkeys**: Cross-platform global hotkey support

## Development

### Project Structure

```
copyman/
├── app.go              # Main application logic
├── main.go             # Application entry point
├── hotkeys.go          # Global hotkey handling
├── wails.json          # Wails configuration
├── go.mod              # Go dependencies
├── frontend/
│   ├── src/
│   │   ├── App.tsx     # Main React component
│   │   ├── App.css     # Styling
│   │   └── main.tsx    # React entry point
│   ├── package.json    # Frontend dependencies
│   └── vite.config.ts  # Vite configuration
└── Makefile           # Build commands
```

### Available Commands

```bash
make dev          # Run in development mode
make build        # Build for current platform
make build-darwin # Build for macOS (universal)
make clean        # Clean build artifacts
make install      # Install dependencies
make test         # Run tests
make package-mac  # Package for macOS distribution
```

## Dependencies

### Go Dependencies
- `github.com/wailsapp/wails/v2` - Desktop app framework
- `github.com/robotn/gohook` - Global hotkey support
- `golang.design/x/clipboard` - Clipboard operations

### Frontend Dependencies
- `react` - UI framework
- `typescript` - Type safety
- `vite` - Build tool

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Troubleshooting

### Global Hotkeys Not Working

The global hotkey functionality requires proper permissions on macOS:

1. Go to System Preferences → Security & Privacy → Privacy
2. Select "Accessibility" from the left panel
3. Add Copyman to the list of allowed applications

### Build Issues

If you encounter build issues:

1. Ensure you have the latest version of Wails CLI
2. Check that Go and Node.js versions meet requirements
3. Try cleaning and reinstalling dependencies:
   ```bash
   make clean
   make install
   ```

### Performance Issues

For optimal performance:
- Keep text bindings reasonably short
- Use the app in overlay mode rather than keeping it always visible
- Consider using the direct hotkey shortcuts (Cmd+Shift+1-9) for fastest access

## Support

For issues and feature requests, please use the [GitHub Issues](https://github.com/vaasu-dhand/copyman/issues) page.