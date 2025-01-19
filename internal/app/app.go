package app

import (
    "mycliai/internal/ai"
    "mycliai/internal/db"
    "mycliai/internal/ui"
)

type App struct {
    config *Config
    db     db.Database
    ai     ai.Client
    ui     *ui.UI
}

func New(cfg *Config, db db.Database) *App {
    return &App{
        config: cfg,
        db:     db,
        ai:     ai.NewGeminiClient(cfg.APIKey),
        ui:     ui.New(),
    }
}

func (a *App) Run() error {
    // Initialize UI and start event loop
    return a.ui.Start()
}
