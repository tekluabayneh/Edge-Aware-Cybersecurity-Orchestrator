package utils

import "github.com/charmbracelet/lipgloss"

// Global styles
var (
	PromptBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(0, 2).
			Foreground(lipgloss.Color("33"))

	ErrorBox = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("196"))

	SuccessBox = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("10")).
			Padding(0, 1).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("34"))
)
