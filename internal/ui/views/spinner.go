package views

import (
    "github.com/charmbracelet/bubbles/spinner"
    "github.com/charmbracelet/lipgloss"
    "mycliai/internal/ui/styles"
)

func RenderSpinner(s spinner.Model) string {
    spinnerLine := lipgloss.JoinHorizontal(lipgloss.Left,
        s.View(),
        " Thinking...")
    return styles.Base.Render(spinnerLine) + "\n"
}
