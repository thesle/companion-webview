package main

import (
	"embed"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the URL from environment variable
	webviewURL := os.Getenv("WEBVIEW_URL")
	if webviewURL == "" {
		log.Fatal("WEBVIEW_URL not set in .env file")
	}

	// Create application with struct
	app := NewApp(webviewURL)

	// Create menu with keyboard shortcuts
	appMenu := menu.NewMenu()
	fileMenu := appMenu.AddSubmenu("Companion WebView")
	fileMenu.AddText("Close", keys.CmdOrCtrl("w"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})
	fileMenu.AddText("Maximize/Restore", keys.CmdOrCtrl("m"), func(_ *menu.CallbackData) {
		runtime.WindowToggleMaximise(app.ctx)
	})
	fileMenu.AddText("Minimize", keys.Combo("m", keys.ControlKey, keys.ShiftKey), func(_ *menu.CallbackData) {
		runtime.WindowMinimise(app.ctx)
	})

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Companion WebView",
		Width:  1000,
		Height: 560,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup:        app.startup,
		Menu:             appMenu,
		Bind: []interface{}{
			app,
		},
		// Frameless window for minimal chrome
		Frameless: true,
		// Linux specific options
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
		},
		// Windows specific options (in case you want to test on Windows)
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
		// Mac specific options (in case you want to test on Mac)
		Mac: &mac.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
