package ui

import (
    "github.com/charmbracelet/bubbles/spinner"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/glamour"
)

func (ui *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        return ui.handleKeyMsg(msg)
    case tea.WindowSizeMsg:
        return ui.handleWindowSize(msg)
    case AIResponseMsg:
        return ui.handleAIResponse(msg)
    case spinner.TickMsg:
        return ui.handleSpinnerTick(msg)
    }

    return ui, cmd
}

func (ui *UI) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    if ui.loading {
        // Only process control characters when loading
        if msg.Type == tea.KeyCtrlC {
            return ui, tea.Quit
        }
        return ui, nil
    }

    switch msg.Type {
    case tea.KeyCtrlC:
        return ui, tea.Quit
    case tea.KeyEnter:
        if ui.input != "" {
            ui.loading = true
            ui.messages = append(ui.messages, Message{Role: "user", Content: ui.input})
            ui.input = ""
            return ui, tea.Batch(
                ui.spinner.Tick,
                ui.getAIResponse,
            )
        }
    case tea.KeyBackspace, tea.KeyDelete:
        if len(ui.input) > 0 {
            ui.input = ui.input[:len(ui.input)-1]
        }
    case tea.KeySpace:
        ui.input += " "
    default:
        // Only add printable ASCII characters
        if len(msg.Runes) > 0 {
            r := msg.Runes[0]
            if r >= 32 && r <= 126 { // ASCII printable range
                ui.input += string(r)
            }
        }
    }

    return ui, nil
}

func (ui *UI) handleWindowSize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
    ui.width = msg.Width
    // Update glamour renderer with new width
    ui.renderer, _ = glamour.NewTermRenderer(
        glamour.WithAutoStyle(),
        glamour.WithWordWrap(msg.Width-2), // Leave some margin
    )
    return ui, nil
}

func (ui *UI) handleAIResponse(msg AIResponseMsg) (tea.Model, tea.Cmd) {
    ui.loading = false
    ui.messages = append(ui.messages, Message{Role: "assistant", Content: string(msg)})
    return ui, nil
}

func (ui *UI) handleSpinnerTick(msg spinner.TickMsg) (tea.Model, tea.Cmd) {
    if ui.loading {
        var spinCmd tea.Cmd
        ui.spinner, spinCmd = ui.spinner.Update(msg)
        return ui, spinCmd
    }
    return ui, nil
}
