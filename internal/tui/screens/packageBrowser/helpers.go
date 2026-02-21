package packagebrowser

import (
	"strings"

	"github.com/ElitistNoob/pacdude/internal/backend"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *PackageBrowserModel) setListItems(output []byte) tea.Cmd {
	return func() tea.Msg {
		o := m.backend.ParseOutput(output)
		items := make([]list.Item, len(o))
		for i, v := range o {
			items[i] = v
		}
		return m.tabs.Active().SetItems(items)
	}
}

func (m *PackageBrowserModel) getSelectedPackage() string {
	var pkg string
	selectedPkg := m.tabs.Active().SelectedItem()
	p, ok := selectedPkg.(backend.Pkg)
	if ok {
		pkg = strings.Split(p.Name, " ")[0]
	}
	return pkg
}

func (m *PackageBrowserModel) isCurrentListEmpty() bool {
	return len(m.tabs.Tabs[m.tabs.Index].Items()) == 0
}

func (m *PackageBrowserModel) loadPackageData(cmd tea.Cmd) tea.Cmd {
	if m.isCurrentListEmpty() {
		var cmds tea.BatchMsg
		cmds = append(cmds, m.tabs.Active().ToggleSpinner())
		cmds = append(cmds, func() tea.Msg { return cmd() })
		return tea.Batch(cmds...)
	}
	return nil
}
