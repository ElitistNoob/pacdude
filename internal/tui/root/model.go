package root

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cmds   []string
	cursor int
}

func InitialModel() model {
	return model{
		cmds: []string{
			"Update Packages",
			"Installed Packages",
		},
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
