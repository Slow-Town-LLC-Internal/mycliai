package ui

import (
    "github.com/charmbracelet/lipgloss"
)

// Styles is our exported styles collection
var Styles = struct {
    You     lipgloss.Style
    AI      lipgloss.Style
    Subtle  lipgloss.Style
    Spinner lipgloss.Style
}{
    You:     lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
    AI:      lipgloss.NewStyle().Foreground(lipgloss.Color("2")),
    Subtle:  lipgloss.NewStyle().Foreground(lipgloss.Color("241")),
    Spinner: lipgloss.NewStyle().Foreground(lipgloss.Color("69")),
}
