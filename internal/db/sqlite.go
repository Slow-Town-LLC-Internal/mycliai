package db

import (
    "context"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
    db *sql.DB
}

func NewSQLite(path string) (*SQLiteDB, error) {
    db, err := sql.Open("sqlite3", path)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    if err := initSchema(db); err != nil {
        return nil, err
    }

    return &SQLiteDB{db: db}, nil
}

func initSchema(db *sql.DB) error {
    schema := `
    CREATE TABLE IF NOT EXISTS sessions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS messages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        session_id INTEGER,
        role TEXT NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY(session_id) REFERENCES sessions(id)
    );`

    _, err := db.Exec(schema)
    return err
}
