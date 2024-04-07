package kuddle

import (
	"context"
	"os"
	"path/filepath"
)

// App struct
type App struct {
	ctx      context.Context
	sessions map[string]Session
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) Startup(ctx context.Context) {
	app.ctx = ctx
}

var kuddleDir string

func GetKuddleDir() string {
	if kuddleDir != "" {
		return kuddleDir
	}

	home, err := os.UserHomeDir()
	if err != nil {
		panic("No home directory found.")
	}
	kuddleDir = filepath.Join(home, ".kuddle")

	return kuddleDir
}
