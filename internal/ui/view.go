package ui

import (
    "strings"
    "github.com/muesli/reflow/wordwrap"
)

func (ui *UI) View() string {
    var sb strings.Builder
    contentWidth := ui.width
    if contentWidth == 0 {
        contentWidth = 80 // fallback width
    }

    // Add initial padding
    sb.WriteString("  ") // Add 2-space padding at start

    // Render messages with proper wrapping
    for _, msg := range ui.messages {
        switch msg.Role {
        case "user":
            prefix := Styles.You.Render("You: ")
            // Add 4 to account for left padding and safety margin
            wrapped := wordwrap.String(msg.Content, contentWidth-len([]rune(prefix))-4)
            lines := strings.Split(wrapped, "\n")
            for i, line := range lines {
                if i == 0 {
                    sb.WriteString(prefix + line + "\n")
                } else {
                    sb.WriteString(strings.Repeat(" ", len([]rune(prefix))) + line + "\n")
                }
                if i < len(lines)-1 {
                    sb.WriteString("  ") // Add padding for wrapped lines
                }
            }
        case "assistant":
            prefix := Styles.AI.Render("AI: ")
            rendered, err := ui.renderer.Render(msg.Content)
            if err != nil {
                rendered = wordwrap.String(msg.Content, contentWidth-len([]rune(prefix))-4)
            }
            rendered = strings.TrimSpace(rendered) // Remove extra newlines from glamour
            lines := strings.Split(rendered, "\n")
            for i, line := range lines {
                if i == 0 {
                    sb.WriteString(prefix + line + "\n")
                } else {
                    sb.WriteString(strings.Repeat(" ", len([]rune(prefix))) + line + "\n")
                }
                if i < len(lines)-1 {
                    sb.WriteString("  ") // Add padding for wrapped lines
                }
            }
        default:
            wrapped := wordwrap.String(msg.Content, contentWidth-4)
            lines := strings.Split(wrapped, "\n")
            for _, line := range lines {
                sb.WriteString("  " + line + "\n")
            }
        }
    }

    // Show spinner when loading
    if ui.loading {
        sb.WriteString("  " + ui.spinner.View())
        sb.WriteString(" Thinking...\n")
    }

    // Input prompt with cursor
    prompt := "> "
    inputWidth := contentWidth - len(prompt) - 4 // Account for padding
    if inputWidth < 0 {
        inputWidth = 0
    }
    wrapped := wordwrap.String(ui.input, inputWidth)
    lines := strings.Split(wrapped, "\n")
    
    sb.WriteString("\n  ") // Add padding before input
    for i, line := range lines {
        if i == 0 {
            sb.WriteString(prompt + line)
            if !ui.loading && ui.cursor && len(lines) == 1 {
                sb.WriteString("█")
            }
            sb.WriteString("\n")
        } else {
            sb.WriteString("  " + strings.Repeat(" ", len(prompt)) + line + "\n")
        }
    }
    
    sb.WriteString("\n  ") // Add padding before quit message
    sb.WriteString(Styles.Subtle.Render("Press Ctrl+D to quit"))

    return sb.String()
}
