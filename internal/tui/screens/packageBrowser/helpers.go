package packagebrowser

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *PackageBrowserModel) setListItems(output []byte) tea.Cmd {
	return func() tea.Msg {
		o := m.Backend.ParseOutput(output)
		items := make([]list.Item, len(o))
		for i, v := range o {
			items[i] = v
		}
		return m.tabContent[m.activeTab].SetItems(items)
	}
}
