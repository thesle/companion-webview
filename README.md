# Companion WebView

A minimal Go application for Linux Mint that displays a web page in a frameless window with minimal chrome, designed for touchscreen displays.

## Features

- **Frameless Window**: No window decorations, maximizing screen real estate
- **Keyboard Shortcuts**: Control the window without mouse/touch
  - `Ctrl+W`: Close the window
  - `Ctrl+M`: Toggle maximize/restore window
  - `Ctrl+Shift+M`: Minimize window
- **Configurable URL**: Set the URL via `.env` file
- **Hardware Acceleration**: Uses GPU acceleration for smooth rendering

## Prerequisites

### System Dependencies (Linux Mint/Ubuntu/Debian)

```bash
sudo apt update
sudo apt install -y build-essential pkg-config libgtk-3-dev libwebkit2gtk-4.0-dev
```

### Go Installation

Make sure you have Go 1.21 or later installed:

```bash
go version
```

If not installed, download from [golang.org](https://golang.org/dl/)

### Wails CLI

Install the Wails CLI tool:

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

Make sure `~/go/bin` is in your PATH:

```bash
export PATH=$PATH:~/go/bin
```

## Installation

1. Clone or navigate to this directory:
```bash
cd /home/conrad/coding/companion-webview
```

2. Install Go dependencies:
```bash
go mod download
```

3. (Optional) Create a custom icon:
```bash
./create-icon.sh
```
Or place your own 256x256 PNG icon at `build/appicon.png`

## Configuration

Edit the `.env` file to set your desired URL:

```bash
WEBVIEW_URL=http://127.0.0.1:8000/tablet
```

## Running the Application

### Development Mode

Run in development mode with hot reload:

```bash
wails dev
```

### Production Build

Build the application:

```bash
wails build
```

The compiled binary will be in `build/bin/companion-webview`

Run the built application:

```bash
./build/bin/companion-webview
```

## Alternative: Simple Go Build

If you prefer not to use the Wails CLI, you can build with standard Go tools:

```bash
go build -o companion-webview .
./companion-webview
```

## Usage

1. Ensure your companion web server is running at the URL specified in `.env`
2. Launch the application
3. The web view will load in a frameless window
4. Use keyboard shortcuts to control the window:
   - `Ctrl+W` to close
   - `Ctrl+M` to maximize/restore
   - `Ctrl+Shift+M` to minimize

## Touchscreen Optimization

For touchscreen displays, you may want to:

1. **Auto-start on boot**: Add the application to your startup applications
2. **Hide cursor**: Use `unclutter` to hide the mouse cursor:
   ```bash
   sudo apt install unclutter
   unclutter -idle 0.1 &
   ```
3. **Prevent screen blanking**: Disable screen saver in system settings

## Project Structure

```
companion-webview/
├── .env                    # Configuration file (URL)
├── .env.example           # Example configuration
├── main.go                # Application entry point
├── app.go                 # Application logic and methods
├── go.mod                 # Go module dependencies
├── Makefile               # Build automation
├── create-icon.sh         # Icon generation script
├── build/
│   └── appicon.png        # Application icon (256x256 PNG)
├── frontend/
│   └── dist/
│       ├── index.html     # HTML wrapper for iframe
│       └── app.js         # JavaScript for keyboard shortcuts
└── README.md              # This file
```

## Troubleshooting

### Dependencies not found

Run:
```bash
go mod tidy
go mod download
```

### WebKit errors on Linux

Ensure WebKit2GTK is installed:
```bash
sudo apt install libwebkit2gtk-4.0-dev
```

### Window doesn't appear

Check that your URL in `.env` is accessible:
```bash
curl http://127.0.0.1:8000/tablet
```

## License

MIT
