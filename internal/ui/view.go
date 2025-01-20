package ui

import (
    "github.com/charmbracelet/lipgloss"
)

// Base styles
var (
    baseStyle = lipgloss.NewStyle().
        PaddingLeft(2).
        PaddingRight(2)

    messageStyle = baseStyle.Copy()

    inputStyle = baseStyle.Copy().
        PaddingTop(1).
        PaddingBottom(1)

    promptStyle = lipgloss.NewStyle().
        PaddingLeft(2).
        Bold(true)

    footerStyle = baseStyle.Copy().
        PaddingTop(1).
        Foreground(lipgloss.Color("241"))
)

func (ui *UI) View() string {
    var sections []string

    // Render messages
    for _, msg := range ui.messages {
        var messageContent string
        switch msg.Role {
        case "user":
            prefix := Styles.You.Render("You:")
            if rendered, err := ui.renderer.Render(msg.Content); err == nil {
                messageContent = rendered
            } else {
                messageContent = msg.Content
            }
            sections = append(sections, lipgloss.JoinVertical(lipgloss.Left,
                promptStyle.Render(prefix),
                messageStyle.Render(messageContent)))

        case "assistant":
            prefix := Styles.AI.Render("AI:")
            if rendered, err := ui.renderer.Render(msg.Content); err == nil {
                messageContent = rendered
            } else {
                messageContent = msg.Content
            }
            sections = append(sections, lipgloss.JoinVertical(lipgloss.Left,
                promptStyle.Render(prefix),
                messageStyle.Render(messageContent)))

        default:
            sections = append(sections, messageStyle.Render(msg.Content))
        }
    }

    // Show spinner when loading
    if ui.loading {
        spinnerLine := lipgloss.JoinHorizontal(lipgloss.Left,
            ui.spinner.View(),
            " Thinking...")
        sections = append(sections, messageStyle.Render(spinnerLine))
    }

    // Input area
    inputLine := lipgloss.JoinHorizontal(lipgloss.Left,
        promptStyle.Render(">"),
        ui.input)
    if !ui.loading && ui.cursor {
        inputLine += "█"
    }
    sections = append(sections, inputStyle.Render(inputLine))

    // Footer
    sections = append(sections, footerStyle.Render("Press Ctrl+D to quit"))

    // Join all sections
    return lipgloss.JoinVertical(lipgloss.Left,
        sections...,
    )
}
