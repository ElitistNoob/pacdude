package backendselector

import (
	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	pb "github.com/ElitistNoob/pacdude/internal/tui/screens/packageBrowser"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *backendSelectorModel) Update(msg tea.Msg) (app.Screen, tea.Cmd) {
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
			var b backend.BackendInterface

			switch selectedChoice {
			case "Pacman":
				b = backend.PacmanBackend{}
				newScreen := pb.NewModel(b)
				return m, func() tea.Msg { return app.ChangeScreenMsg{NewScreen: newScreen} }
			}
		}
	}

	return m, nil
}
