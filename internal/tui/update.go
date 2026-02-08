package tui

import (
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
	"github.com/ElitistNoob/pacdude/internal/tui/pkgs"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		updated, cmd := m.current.Update(msg)
		m.current = updated
		return m, cmd

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "backspace":
			if m.previous != nil {
				m.current, m.previous = m.previous, m.current
			}
			return m, nil
		}

		updated, cmd := m.current.Update(msg)
		m.current = updated
		return m, cmd

	case messages.GoToPkgs:
		m.previous = m.current
		p := pkgs.NewPkgsModel()

		updatedModel, _ := p.Update(tea.WindowSizeMsg{
			Width:  m.width,
			Height: m.height,
		})
		m.current = updatedModel

		return m, pkgs.ExecWrapper(msg.Args)
	}

	updated, cmd := m.current.Update(msg)
	m.current = updated
	return m, cmd
}
