package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var CursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("10"))

var TabActive = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12"))

var TabInactive = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#fff"))

var DocStyle = lipgloss.NewStyle().Margin(1, 2)

var ListStyle = lipgloss.NewStyle().
	Padding(1, 1).
	Border(lipgloss.RoundedBorder(), false, true, true, true)
