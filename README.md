# MyCliAI

A simple terminal-based chat interface for Google's Gemini AI model, built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea). Features markdown rendering and local conversation storage.

## Prerequisites

- Go 1.21 or later
- Google Gemini API key ([Get one here](https://makersuite.google.com/app/apikey))

## Quick Start

1. Clone and build:
```bash
git clone https://github.com/yourusername/mycliai.git
cd mycliai
go mod tidy
```

2. Set your API key:
```bash
export GEMINI_API=your_api_key_here
```

3. Run:
```bash
go run cmd/mycliai/main.go
```

## Usage

- Type your message and press Enter to send
- Press Ctrl+C to quit
- Responses support markdown formatting
- Local SQLite storage for conversation history

## Known Issues

- Initial input may show control characters
- Some markdown formatting issues in responses
