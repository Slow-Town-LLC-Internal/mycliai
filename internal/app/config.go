package app

import (
    "os"
    "path/filepath"
)

type Config struct {
    APIKey  string
    DBPath  string
    DataDir string
}

func LoadConfig() (*Config, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return nil, err
    }

    dataDir := filepath.Join(homeDir, ".mycliai")
    if err := os.MkdirAll(dataDir, 0755); err != nil {
        return nil, err
    }

    return &Config{
        APIKey:  os.Getenv("GEMINI_API"),
        DBPath:  filepath.Join(dataDir, "mycliai.db"),
        DataDir: dataDir,
    }, nil
}
