package tui

import (
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
	"github.com/ElitistNoob/pacdude/internal/tui/pkgs"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Window Resize Messages
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		updated, cmd := m.current.Update(msg)
		m.current = updated
		return m, cmd

	// Keypress Messages
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "backspace":
			if m.previous != nil {
				m.current, m.previous = m.previous, m.current
			}
			return m, nil
		case "y":
			m.state = stateInstalling
			return m, m.installPkgCmd(m.selectedPkg)
		case "n":
			m.state = stateReady
			return m, nil
		}

		updated, cmd := m.current.Update(msg)
		m.current = updated
		return m, cmd

	// Event Trigger Messages
	case messages.GoToPkgsMsg:
		m.previous = m.current
		p := pkgs.NewPkgsModel()

		updatedModel, _ := p.Update(tea.WindowSizeMsg{
			Width:  m.width,
			Height: m.height,
		})
		m.current = updatedModel
		return m, pkgs.ExecWrapper(&pkgs.PacmanOpts{}, msg.Args)
	case messages.InstallPkgMsg:
		m.state = stateConfirmInstall
		m.selectedPkg = msg.Args[0]
		return m, nil
	case messages.InstallResultMsg:
		if msg.Err.Err != nil {
			m.state = stateError
			m.lastErr.err = msg.Err.Err
			return m, nil
		}
		m.state = stateDone
		m.lastOutput = outputMsg(msg.Output)
		return m, nil
	}

	updated, cmd := m.current.Update(msg)
	m.current = updated
	return m, cmd
}
