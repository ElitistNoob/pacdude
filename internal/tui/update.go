package tui

import (
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
	"github.com/ElitistNoob/pacdude/internal/tui/pkgs"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.previous != nil {
				m.current, m.previous = m.previous, m.current
			}
			return m, nil
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
		}

		updated, cmd := m.current.Update(msg)
		m.current = updated
		return m, cmd

	case messages.GoToPkgs:
		m.previous = m.current
		m.current = pkgs.NewPkgsModel()
		return m, pkgs.ExecWrapper(msg.Args)

	case messages.PkgOutput:
		updated, cmd := m.current.Update(msg)
		m.current = updated
		return m, cmd
	}

	return m, nil
}
