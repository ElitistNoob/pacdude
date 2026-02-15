package backendselector

import (
	"github.com/ElitistNoob/pacdude/internal/app"
	tea "github.com/charmbracelet/bubbletea"
)

type backendSelectorModel struct {
	choices []string
	cursor  int
}

func NewModel() app.Screen {
	return &backendSelectorModel{
		cursor:  0,
		choices: []string{"Pacman"},
	}
}

func (m *backendSelectorModel) Init() tea.Cmd {
	return nil
}
