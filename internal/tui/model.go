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
	lastOutput outputMsg
	lastErr    errorMsg
	cursor     int
	choices    []string
}

func newTuiModel() *model {
	return &model{
		current:    choice.InitialChoiceModel(),
		previous:   nil,
		lastOutput: "",
		lastErr:    errorMsg{},
		choices:    []string{},
		cursor:     0,
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}
