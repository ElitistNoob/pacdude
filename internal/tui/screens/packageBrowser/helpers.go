package packagebrowser

import (
	"strings"

	"github.com/ElitistNoob/pacdude/internal/backend"
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

func (m *PackageBrowserModel) getSelectedPackage() string {
	var pkg string
	selectedPkg := m.tabContent[m.activeTab].SelectedItem()
	p, ok := selectedPkg.(backend.Pkg)
	if ok {
		pkg = strings.Split(p.Name, " ")[0]
	}
	return pkg
}

func (m *PackageBrowserModel) isCurrentListEmpty() bool {
	return len(m.tabContent[m.activeTab].Items()) == 0
}

func (m *PackageBrowserModel) loadPackageData(cmd tea.Cmd) tea.Cmd {
	activeList := m.tabContent[m.activeTab]
	if m.isCurrentListEmpty() {
		var cmds tea.BatchMsg
		cmds = append(cmds, activeList.ToggleSpinner())
		cmds = append(cmds, func() tea.Msg { return cmd() })
		return tea.Batch(cmds...)
	}
	return nil
}
