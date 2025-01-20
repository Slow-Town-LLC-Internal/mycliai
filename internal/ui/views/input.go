package views

import (
    "strings"
    "github.com/charmbracelet/lipgloss"
    "mycliai/internal/ui/styles"
)

func RenderInput(input string, width int, showCursor bool, loading bool) string {
    var sb strings.Builder
    inputWidth := width - 16

    if len(input) > 0 {
        wrapped := lipgloss.WrapSoft(input, inputWidth)
        prompt := styles.Prompt.Render(">  ")
        sb.WriteString("\n")
        sb.WriteString(prompt)
        sb.WriteString(styles.Content.Render(wrapped))
        if !loading && showCursor {
            sb.WriteString("█")
        }
    } else {
        prompt := styles.Prompt.Render(">  ")
        sb.WriteString("\n")
        sb.WriteString(prompt)
        if !loading && showCursor {
            sb.WriteString("█")
        }
    }

    return sb.String()
}
