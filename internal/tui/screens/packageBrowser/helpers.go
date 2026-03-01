package packagebrowser

import (
	"strings"

	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	panels "github.com/ElitistNoob/pacdude/internal/tui/components"
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m *PackageBrowserModel) setListItems(pkgs []backend.Pkg) tea.Cmd {
	return func() tea.Msg {
		items := make([]list.Item, len(pkgs))

		for i, v := range pkgs {
			items[i] = v
		}

		return m.tabs.Active().(*panels.ListPanel).List.SetItems(items)
	}
}

func (m *PackageBrowserModel) getSelectedPackage() string {
	item := m.tabs.Active().(*panels.ListPanel).SelectedItem()

	p, ok := item.(backend.Pkg)
	if !ok {
		return ""
	}

	return strings.Split(p.Name, " ")[0]
}

func runBackend(fn func() backend.ResultMsg) tea.Cmd {
	return func() tea.Msg {
		res := fn()

		if res.Err.Err != nil {
			return messages.ActionMsg{
				Type: res.ActionType,
				Err:  res.Err.Err,
			}
		}

		return messages.ActionMsg{
			Type:    res.ActionType,
			Payload: res.Output,
		}
	}
}

func (m *PackageBrowserModel) getContentSize(style lipgloss.Style) (int, int) {
	w, h := style.GetFrameSize()
	return m.width - w, m.height - h
}

func (m *PackageBrowserModel) loadActive() tea.Cmd {
	panel := m.tabs.Active()

	lp, ok := panel.(*panels.ListPanel)
	if !ok {
		return nil
	}

	if lp.IsEmpty() {
		return nil
	}

	switch m.tabs.Index {
	case 0:
		return runBackend(m.backend.ListAll)
	case 1:
		return runBackend(m.backend.ListInstalled)
	case 2:
		return runBackend(m.backend.ListUpgradable)
	}

	return nil
}

func (m *PackageBrowserModel) setBackend(index int) tea.Cmd {
	switch index {
	case 0:
		m.backend = backend.PacmanBackend{}
	case 1:
		m.backend = backend.FlatpakBackend{}
	case 2:
		m.backend = backend.BrewBackend{}
	}

	newScreen := NewModel(m.backend, m.managerTab.Index)
	return func() tea.Msg {
		return app.ChangeScreenMsg{
			NewScreen: newScreen,
		}
	}
}
