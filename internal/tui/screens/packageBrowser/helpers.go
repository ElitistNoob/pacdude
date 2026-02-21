package packagebrowser

import (
	"strings"

	"github.com/ElitistNoob/pacdude/internal/backend"
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *PackageBrowserModel) setListItems(pkgs []backend.Pkg) tea.Cmd {
	return func() tea.Msg {
		items := make([]list.Item, len(pkgs))

		for i, v := range pkgs {
			items[i] = v
		}

		return m.tabs.Active().SetItems(items)
	}
}

func (m *PackageBrowserModel) getSelectedPackage() string {
	item := m.tabs.SelectedItem()

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
