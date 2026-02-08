package choice

import (
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var args []string

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			selectedChoice := m.choices[m.cursor]
			switch selectedChoice {
			case "Installed Packages":
				args = []string{"-Qs"}
			}
			return m, messages.MsgHandler(args)
		}
	}

	return m, nil
}
