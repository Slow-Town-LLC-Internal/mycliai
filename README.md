# creating a cli ai interface with Claude to be used in my shell. 

## directory structure

```
project/
├── cmd/
│   └── mycliai/
│       └── main.go
├── internal/
│   ├── app/
│   │   ├── app.go        // Main application logic
│   │   └── config.go     // Configuration handling
│   ├── chat/
│   │   ├── chat.go       // Chat session management
│   │   ├── message.go    // Message types and handling
│   │   └── history.go    // Conversation history
│   ├── db/
│   │   ├── db.go         // Database interface
│   │   ├── schema.go     // Database schema
│   │   └── sqlite.go     // SQLite implementation
│   ├── ui/
│   │   ├── ui.go         // TUI components
│   │   ├── styles.go     // UI styles
│   │   └── components.go // Reusable UI components
│   └── ai/
│       ├── client.go     // AI client interface
│       └── gemini.go     // Gemini implementation
├── pkg/
│   └── config/
│       └── config.go     // Public config package
├── go.mod
└── README.md

```
