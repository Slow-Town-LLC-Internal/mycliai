package ui

import (
    "strings"
    "github.com/muesli/reflow/wordwrap"
    "mycliai/internal/ui/styles"
    "mycliai/internal/ui/views"
)

func (ui *UI) View() string {
    var sb strings.Builder
    width := ui.width
    if width == 0 {
        width = 80
    }

    // Account for terminal padding and styling characters
    contentWidth := width - 24  // Increased padding margin

    // Render messages
    for _, msg := range ui.messages {
        switch msg.Role {
        case "user", "assistant":
            sb.WriteString(views.RenderMessage(msg.Role, msg.Content, contentWidth, ui.renderer))
        default:
            wrapped := wordwrap.String(msg.Content, contentWidth)
            sb.WriteString(styles.Base.Render(wrapped))
            sb.WriteString("\n")
        }
    }

    // Show spinner when loading
    if ui.loading {
        sb.WriteString(views.RenderSpinner(ui.spinner))
    }

    // Input area
    sb.WriteString(views.RenderInput(ui.input, contentWidth, ui.cursor, ui.loading))

    // Footer
    sb.WriteString("\n\n")
    sb.WriteString(styles.Footer.Render("Press Ctrl+D to quit"))

    return sb.String()
}
