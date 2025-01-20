package ui

import (
    "github.com/charmbracelet/bubbles/spinner"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/glamour"
)

func (ui *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    // Handle window size updates immediately
    if wmsg, ok := msg.(tea.WindowSizeMsg); ok {
        if !ui.ready {
            ui.ready = true
        }
        ui.width = wmsg.Width
        ui.renderer, _ = glamour.NewTermRenderer(
            glamour.WithAutoStyle(),
            glamour.WithWordWrap(wmsg.Width-2),
        )
        cmds = append(cmds, nil)
    }

    // Skip other updates until ready
    if !ui.ready {
        return ui, nil
    }

    switch msg := msg.(type) {
    case tea.KeyMsg:
        // Quick exit if loading
        if ui.loading {
            if msg.Type == tea.KeyCtrlD {
                return ui, tea.Quit
            }
            return ui, nil
        }

        // Handle input
        switch msg.Type {
        case tea.KeyCtrlD:
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

    case AIResponseMsg:
        ui.loading = false
        ui.messages = append(ui.messages, Message{Role: "assistant", Content: string(msg)})
        cmds = append(cmds, nil)

    case spinner.TickMsg:
        if ui.loading {
            var spinCmd tea.Cmd
            ui.spinner, spinCmd = ui.spinner.Update(msg)
            cmds = append(cmds, spinCmd)
        }
    }

    return ui, tea.Batch(cmds...)
}
