package db

import (
    "context"
    "database/sql"
    "time"
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

// Implementation of Database interface methods
func (db *SQLiteDB) CreateSession(ctx context.Context) (int64, error) {
    result, err := db.db.ExecContext(ctx, "INSERT INTO sessions DEFAULT VALUES")
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func (db *SQLiteDB) SaveMessage(ctx context.Context, msg *Message) error {
    query := `
        INSERT INTO messages (session_id, role, content, created_at)
        VALUES (?, ?, ?, ?)`
    
    _, err := db.db.ExecContext(ctx, query, msg.SessionID, msg.Role, msg.Content, time.Now())
    return err
}

func (db *SQLiteDB) GetSessionMessages(ctx context.Context, sessionID int64) ([]Message, error) {
    query := `
        SELECT id, session_id, role, content, created_at
        FROM messages
        WHERE session_id = ?
        ORDER BY created_at ASC`
    
    rows, err := db.db.QueryContext(ctx, query, sessionID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var messages []Message
    for rows.Next() {
        var msg Message
        err := rows.Scan(&msg.ID, &msg.SessionID, &msg.Role, &msg.Content, &msg.CreatedAt)
        if err != nil {
            return nil, err
        }
        messages = append(messages, msg)
    }
    
    return messages, rows.Err()
}

func (db *SQLiteDB) Close() error {
    if db.db != nil {
        return db.db.Close()
    }
    return nil
}
