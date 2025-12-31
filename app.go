package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx        context.Context
	webviewURL string
}

// NewApp creates a new App application struct
func NewApp(url string) *App {
	return &App{
		webviewURL: url,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetURL returns the URL to load in the webview
func (a *App) GetURL() string {
	return a.webviewURL
}

// CloseWindow closes the application window
func (a *App) CloseWindow() {
	runtime.Quit(a.ctx)
}

// MaximizeWindow toggles window maximization
func (a *App) MaximizeWindow() {
	runtime.WindowToggleMaximise(a.ctx)
}

// MinimizeWindow minimizes the window
func (a *App) MinimizeWindow() {
	runtime.WindowMinimise(a.ctx)
}
