package styles

import (
    "github.com/charmbracelet/lipgloss"
)

var (
    Base = lipgloss.NewStyle().
        PaddingLeft(2).
        PaddingRight(4)

    Content = lipgloss.NewStyle().
        PaddingLeft(4)

    Prompt = lipgloss.NewStyle().
        PaddingLeft(2).
        Width(6).
        Bold(true)

    Footer = lipgloss.NewStyle().
        PaddingTop(1).
        PaddingBottom(1).
        Foreground(lipgloss.Color("241"))

    You = lipgloss.NewStyle().
        Foreground(lipgloss.Color("5"))

    AI = lipgloss.NewStyle().
        Foreground(lipgloss.Color("2"))

    Spinner = lipgloss.NewStyle().
        Foreground(lipgloss.Color("69"))
)
