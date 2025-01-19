package ui

import (
    tea "github.com/charmbracelet/bubbletea"
)

type UI struct{}

func New() *UI {
    return &UI{}
}

func (ui *UI) Init() tea.Cmd {
    return nil
}

func (ui *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.Type == tea.KeyCtrlC {
            return ui, tea.Quit
        }
    }
    return ui, nil
}

func (ui *UI) View() string {
    return "Welcome to MyCliAI!\nPress Ctrl+C to quit\n"
}

func (ui *UI) Start() error {
    program := tea.NewProgram(ui)
    _, err := program.Run()
    return err
}
