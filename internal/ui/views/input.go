package views

import (
    "strings"
    "github.com/muesli/reflow/wordwrap"
    "mycliai/internal/ui/styles"
)

func RenderInput(input string, width int, showCursor bool, loading bool) string {
    var sb strings.Builder

    if len(input) > 0 {
        wrapped := wordwrap.String(input, width)
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
