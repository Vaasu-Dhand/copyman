# Copyman

A lightweight, fast clipboard manager for macOS that puts your most-used text snippets at your fingertips.

![Copyman Screenshot](frontend/public/favicon-light.png)

## Features

- 🚀 **Lightning Fast**: Access your clipboard with `Ctrl+Shift+Space`
- 🔢 **Quick Copy**: Bind text to number keys 1-9 for instant copying
- 🎨 **Beautiful Interface**: Clean, modern design with dark/light themes
- ⌨️ **Global Hotkeys**: Works from anywhere on your Mac
- 💾 **Persistent Storage**: Your text snippets are saved between sessions
- 🪶 **Lightweight**: Minimal resource usage, stays out of your way

## Installation

### Download Pre-built App (Recommended)

1. Go to the [Releases](https://github.com/yourusername/copyman/releases) page
2. Download the latest `Copyman-darwin-universal.zip`
3. Unzip and drag `Copyman.app` to your Applications folder
4. Right-click the app and select "Open" (required for first launch due to macOS security)

### Grant Permissions

Copyman needs accessibility permissions to work with global hotkeys:

1. Open **System Preferences** → **Security & Privacy** → **Privacy** → **Accessibility**
2. Click the lock icon and enter your password
3. Click the **+** button and add `Copyman.app`
4. Make sure the checkbox next to Copyman is checked

## Usage

### Basic Workflow

1. **Show Copyman**: Press `Ctrl+Shift+Space`
2. **Set up text snippets**: Click ⚙️ Settings and bind text to number keys 1-9
3. **Quick copy**: Press `Ctrl+Shift+1-9` to instantly copy and paste your snippets

### Hotkeys

- `Ctrl+Shift+Space` - Toggle Copyman overlay
- `Ctrl+Shift+1-9` - Quick copy bound text snippets
- `Esc` - Close overlay (when focused)

### Examples of What to Store

- Email signatures
- Common responses
- Code snippets
- Addresses and phone numbers
- URLs you use frequently
- Emoji combinations
- Meeting links

## Building from Source

### Prerequisites

- Go 1.19 or later
- Node.js 16 or later
- Wails CLI: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`

### Development

```bash
# Clone the repository
git clone https://github.com/yourusername/copyman.git
cd copyman

# Install dependencies
wails build

# Run in development mode
wails dev
```

### Building for Distribution

```bash
# Build for macOS
wails build -platform darwin/universal

# The app will be created in build/bin/
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Development Setup

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Troubleshooting

### App won't open
- Make sure you've granted accessibility permissions
- Try right-clicking and selecting "Open" instead of double-clicking

### Hotkeys not working
- Check that Copyman has accessibility permissions in System Preferences
- Make sure no other app is using the same hotkey combinations

### Settings not saving
- Check that you have write permissions to `~/.config/copyman/`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Wails](https://wails.io/) - Go + Web frontend framework
- Icons from [Lucide](https://lucide.dev/)
- Global hotkeys powered by [robotn/gohook](https://github.com/robotn/gohook)

---

**Made with ❤️ for productivity enthusiasts**