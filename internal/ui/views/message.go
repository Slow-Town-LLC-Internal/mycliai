package views

import (
    "strings"
    "github.com/charmbracelet/lipgloss"
    "github.com/muesli/reflow/indents"
    "github.com/charmbracelet/glamour"
    "mycliai/internal/ui/styles"
)

func RenderMessage(role, content string, width int, renderer *glamour.TermRenderer) string {
    var sb strings.Builder
    textWidth := width - 16

    // Render prefix
    prefix := styles.You.Render("You:")
    if role == "assistant" {
        prefix = styles.AI.Render("AI:")
    }

    // Process content
    var rendered string
    if r, err := renderer.Render(content); err == nil {
        rendered = r
    } else {
        // Manual wrapping with proper indentation
        lines := strings.Split(content, "\n")
        for i, line := range lines {
            if strings.TrimSpace(line) != "" {
                wrapped := lipgloss.WrapSoft(line, textWidth)
                if i == 0 {
                    lines[i] = wrapped
                } else {
                    lines[i] = indents.String("    "+wrapped, 4)
                }
            }
        }
        rendered = strings.Join(lines, "\n")
    }

    // Join prefix and content
    sb.WriteString(styles.Prompt.Render(prefix))
    sb.WriteString("\n")
    sb.WriteString(styles.Content.Render(rendered))
    sb.WriteString("\n")

    return sb.String()
}
