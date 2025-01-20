package ui

import (
    "fmt"
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
    // Clear screen and reset cursor before creating UI
    fmt.Print("\033[H\033[2J\033[3J")  // Clear screen and scrollback buffer
    fmt.Print("\033[?25l")             // Hide cursor temporarily

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
    _, err := p.Run()
    return err
}

func (ui *UI) Init() tea.Cmd {
    return tea.Batch(
        tea.EnterAltScreen,
        tea.ClearScreen,
        tea.HideCursor,  // Hide cursor during initialization
        func() tea.Msg {
            // Make sure we're fully initialized before accepting input
            ui.ready = true
            return nil
        },
    )
}
