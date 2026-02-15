package app

import tea "github.com/charmbracelet/bubbletea"

type Screen interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Screen, tea.Cmd)
	View() string
}
