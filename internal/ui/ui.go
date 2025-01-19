package ui

import (
    "github.com/charmbracelet/bubbles/spinner"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/glamour"
    "mycliai/internal/ai"
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
}

type Message struct {
    Role    string
    Content string
}

func New(aiClient ai.Client) *UI {
    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = Styles.Spinner

    renderer, _ := glamour.NewTermRenderer(
        glamour.WithAutoStyle(),
        glamour.WithWordWrap(0), // Will be set dynamically based on terminal width
    )

    return &UI{
        aiClient:  aiClient,
        spinner:   s,
        renderer:  renderer,
        messages:  []Message{{Role: "system", Content: "Welcome to MyCliAI!\nAsk me to write a haiku or anything else!\n"}},
        cursor:    true,
        width:     80, // Default width
    }
}

func (ui *UI) Start() error {
    p := tea.NewProgram(
        ui,
        tea.WithAltScreen(),
    )
    _, err := p.Run()
    return err
}

func (ui *UI) Init() tea.Cmd {
    return tea.Batch(
        ui.spinner.Tick,
        tea.EnterAltScreen,
    )
}
