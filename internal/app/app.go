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
    aiClient := ai.NewGeminiClient(cfg.APIKey)
    return &App{
        config: cfg,
        db:     db,
        ai:     aiClient,
        ui:     ui.New(aiClient),
    }
}

func (a *App) Run() error {
    defer a.ai.Close()
    return a.ui.Start()
}
