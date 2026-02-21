package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var CursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("10"))

var TabActive = lipgloss.NewStyle().
	Background(lipgloss.Color("10")).
	Foreground(lipgloss.Color("#111111")).
	Padding(0, 1)

var TabInactive = lipgloss.NewStyle().
	Background(lipgloss.Color("#111111")).
	Foreground(lipgloss.Color("#fff")).
	Padding(0, 1)

var DocStyle = lipgloss.NewStyle().Margin(1, 2)
