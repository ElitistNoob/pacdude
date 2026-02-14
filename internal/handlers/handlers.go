package handlers

import (
	msg "github.com/ElitistNoob/pacdude/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

func HandleShowPackageScreen(args []string) tea.Cmd {
	return func() tea.Msg {
		return msg.GoToPkgsMsg{Args: args}
	}
}
