package ui

import (
    "strings"
    "github.com/muesli/reflow/wordwrap"
)

func (ui *UI) View() string {
    var sb strings.Builder

    // Render messages with proper wrapping
    for _, msg := range ui.messages {
        switch msg.Role {
        case "user":
            prefix := Styles.You.Render("You: ")
            wrapped := wordwrap.String(msg.Content, ui.width-len([]rune(prefix)))
            sb.WriteString(prefix + wrapped + "\n")
        case "assistant":
            prefix := Styles.AI.Render("AI: ")
            rendered, err := ui.renderer.Render(msg.Content)
            if err != nil {
                rendered = wordwrap.String(msg.Content, ui.width-len([]rune(prefix)))
            }
            sb.WriteString(prefix + rendered + "\n")
        default:
            wrapped := wordwrap.String(msg.Content, ui.width)
            sb.WriteString(wrapped)
        }
    }

    // Show spinner when loading
    if ui.loading {
        sb.WriteString(ui.spinner.View())
        sb.WriteString(" Thinking...\n")
    }

    // Input prompt with cursor
    prompt := "> "
    inputWidth := ui.width - len(prompt)
    wrapped := wordwrap.String(ui.input, inputWidth)
    
    sb.WriteString("\n" + prompt + wrapped)
    if ui.cursor {
        sb.WriteString("█")
    }
    
    sb.WriteString("\n\n")
    sb.WriteString(Styles.Subtle.Render("Press Ctrl+C to quit"))

    return sb.String()
}
