package choice

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor  int
	choices []string
}

func InitialChoiceModel() *model {
	return &model{
		cursor:  0,
		choices: []string{"Installed Packages", "Search Packages"},
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}
