package tui

import (
	"github.com/ElitistNoob/pacdude/internal/tui/choice"
	tea "github.com/charmbracelet/bubbletea"
)

type outputMsg string
type errorMsg struct{ err error }

type model struct {
	current    tea.Model
	previous   tea.Model
	width      int
	height     int
	lastOutput outputMsg
	lastErr    errorMsg
}

func newTuiModel() *model {
	return &model{
		current:    choice.InitialChoiceModel(),
		previous:   nil,
		lastOutput: "",
		lastErr:    errorMsg{},
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}
