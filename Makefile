.PHONY: dev build clean install test

# Development
dev:
	wails dev

# Build for production
build:
	wails build

# Build for distribution (macOS)
build-darwin:
	wails build -platform darwin/universal -clean

# Clean build artifacts
clean:
	rm -rf build/
	rm -rf frontend/dist/
	rm -rf frontend/node_modules/

# Install dependencies
install:
	cd frontend && npm install
	go mod tidy

# Generate Wails bindings
generate:
	wails generate module

# Test
test:
	go test ./...

# Run without development tools
run:
	go run .

# Install global dependencies
deps:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest
	cd frontend && npm install

# Package for macOS
package-mac: build-darwin
	@echo "Creating macOS app bundle..."
	mkdir -p dist
	cp -r build/bin/copyman.app dist/
	@echo "App packaged in dist/ directory"

# Help
help:
	@echo "Available commands:"
	@echo "  dev          - Run in development mode"
	@echo "  build        - Build for current platform"
	@echo "  build-darwin - Build for macOS (Intel + Apple Silicon)"
	@echo "  clean        - Clean build artifacts"
	@echo "  install      - Install dependencies"
	@echo "  generate     - Generate Wails bindings"
	@echo "  test         - Run tests"
	@echo "  run          - Run without dev tools"
	@echo "  deps         - Install global dependencies"
	@echo "  package-mac  - Package for macOS distribution"