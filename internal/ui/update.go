package ui

import (
    "github.com/charmbracelet/bubbles/spinner"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/glamour"
    "strings"
)

func (ui *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    if !ui.ready {
        if _, ok := msg.(tea.WindowSizeMsg); ok {
            ui.ready = true
        }
        return ui, nil
    }

    switch msg := msg.(type) {
    case tea.KeyMsg:
        return ui.handleKeyMsg(msg)
    case tea.WindowSizeMsg:
        return ui.handleWindowSize(msg)
    case AIResponseMsg:
        return ui.handleAIResponse(msg)
    case spinner.TickMsg:
        if ui.loading {
            var spinCmd tea.Cmd
            ui.spinner, spinCmd = ui.spinner.Update(msg)
            return ui, spinCmd
        }
    }

    return ui, nil
}

func (ui *UI) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    // Quick exit for loading state
    if ui.loading {
        if msg.Type == tea.KeyCtrlC {
            return ui, tea.Quit
        }
        return ui, nil
    }

    // Handle special keys first for immediate feedback
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
        
    case tea.KeySpace:
        ui.input += " "
        return ui, tea.Batch() // Force immediate update
        
    case tea.KeyBackspace, tea.KeyDelete:
        if len(ui.input) > 0 {
            ui.input = ui.input[:len(ui.input)-1]
        }
        
    case tea.KeyEsc:
        if ui.input != "" {
            ui.input = ""
        }

    case tea.KeyRunes:
        ui.input += string(msg.Runes)
    }

    return ui, nil
}

func (ui *UI) handleWindowSize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
    if !ui.ready {
        ui.ready = true
    }
    ui.width = msg.Width
    ui.renderer, _ = glamour.NewTermRenderer(
        glamour.WithAutoStyle(),
        glamour.WithWordWrap(msg.Width-2),
    )
    return ui, nil
}

func (ui *UI) handleAIResponse(msg AIResponseMsg) (tea.Model, tea.Cmd) {
    ui.loading = false
    
    response := string(msg)
    if strings.Contains(response, "FinishReasonSafety") {
        response = "I apologize, but I cannot generate that content due to safety filters. Please try rephrasing your request."
    } else if strings.Contains(response, "Error:") {
        response = "An error occurred while processing your request. Please try again."
    }
    
    ui.messages = append(ui.messages, Message{Role: "assistant", Content: response})
    return ui, nil
}
