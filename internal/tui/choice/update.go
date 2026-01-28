package choice

import (
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.choices) - 1
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}

		case "enter", " ":
			selectedChoice := m.choices[m.cursor]
			var args []string
			switch selectedChoice {
			case "Installed Packages":
				args = []string{"-Qs"}
			case "Search Packages":
				args = []string{"-Ss", "steam"}
			}
			return m, messages.MsgHandler(args)
		}
	}

	return m, nil
}
