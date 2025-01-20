package ui

import (
    "os"
    "github.com/charmbracelet/bubbles/spinner"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/glamour"
    "mycliai/internal/ai"
    "mycliai/internal/ui/styles"
)

type UI struct {
    input     string
    cursor    bool
    messages  []Message
    aiClient  ai.Client
    spinner   spinner.Model
    loading   bool
    renderer  *glamour.TermRenderer
    width     int
    ready     bool
}

type Message struct {
    Role    string
    Content string
}

func New(aiClient ai.Client) *UI {
    // Hide cursor and clear screen immediately
    os.Stdout.WriteString("\x1b[?25l")   // Hide cursor
    os.Stdout.WriteString("\x1b[2J")     // Clear screen
    os.Stdout.WriteString("\x1b[H")      // Move to home position
    os.Stdout.Sync()

    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = styles.Spinner

    renderer, _ := glamour.NewTermRenderer(
        glamour.WithAutoStyle(),
        glamour.WithWordWrap(0),
    )

    return &UI{
        aiClient:  aiClient,
        spinner:   s,
        renderer:  renderer,
        messages:  []Message{{Role: "system", Content: "Welcome to MyCliAI!\nAsk me to write a haiku or anything else!\n"}},
        cursor:    true,
        width:     80,
        ready:     false,
    }
}

func (ui *UI) Start() error {
    p := tea.NewProgram(
        ui,
        tea.WithAltScreen(),
        tea.WithInputTTY(),
    )

    // Run the program
    _, err := p.Run()
    return err
}

func (ui *UI) Init() tea.Cmd {
    return tea.Batch(
        tea.EnterAltScreen,
        tea.ClearScreen,
        tea.HideCursor,
        func() tea.Msg {
            ui.ready = true
            return nil
        },
    )
}
