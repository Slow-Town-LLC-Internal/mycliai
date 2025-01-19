package ui

import (
    "github.com/charmbracelet/lipgloss"
)

var Styles = struct {
    You          lipgloss.Style
    AI           lipgloss.Style
    Subtle       lipgloss.Style
    SpinnerColor lipgloss.Color
}{
    You: lipgloss.NewStyle().
        Foreground(lipgloss.Color("5")),
    AI: lipgloss.NewStyle().
        Foreground(lipgloss.Color("2")),
    Subtle: lipgloss.NewStyle().
        Foreground(lipgloss.Color("241")),
    SpinnerColor: lipgloss.Color("69"),
}
