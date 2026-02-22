package packagebrowser

import (
	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *PackageBrowserModel) Update(msg tea.Msg) (app.Screen, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		cmds = append(cmds, m.reduceWindowSize(msg))

	case tea.KeyMsg:
		cmds = append(cmds, m.reduceKeys(msg))

	case messages.ActionMsg:
		cmds = append(cmds, m.reduceActions(msg))
	}

	updated, cmd := m.tabs.Active().Update(msg)
	m.tabs.SetActive(updated)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
