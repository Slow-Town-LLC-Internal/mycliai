package db

import (
    "context"
    "time"
)

type Message struct {
    ID        int64
    SessionID int64
    Role      string
    Content   string
    CreatedAt time.Time
}

type Database interface {
    SaveMessage(ctx context.Context, msg *Message) error
    GetSessionMessages(ctx context.Context, sessionID int64) ([]Message, error)
    CreateSession(ctx context.Context) (int64, error)
    Close() error
}
