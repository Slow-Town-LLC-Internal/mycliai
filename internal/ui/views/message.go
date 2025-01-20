package views

import (
    "strings"
    "github.com/charmbracelet/glamour"
    "github.com/muesli/reflow/wordwrap"
    "mycliai/internal/ui/styles"
)

func RenderMessage(role, content string, width int, renderer *glamour.TermRenderer) string {
    var sb strings.Builder

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
                wrapped := wordwrap.String(line, width)
                if i > 0 {
                    // Add indentation for continuation lines
                    wrapped = strings.ReplaceAll(wrapped, "\n", "\n    ")
                }
                lines[i] = wrapped
            }
        }
        rendered = strings.Join(lines, "\n")
    }

    // Join prefix and content with proper style and margin
    sb.WriteString(styles.Prompt.Render(prefix))
    sb.WriteString("\n")
    sb.WriteString(styles.Content.Render(rendered))
    sb.WriteString("\n")

    return sb.String()
}
