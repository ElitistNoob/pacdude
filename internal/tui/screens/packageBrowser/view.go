package packagebrowser

import (
	"fmt"
	"strings"

	"github.com/ElitistNoob/pacdude/internal/backend"
)

func (m *PackageBrowserModel) View() string {
	var b strings.Builder
	for i, tab := range m.tabContent {
		if i == m.activeTab {
			b.WriteString(fmt.Sprintf("[%s] ", tab.Title))
		} else {
			b.WriteString(fmt.Sprintf("%s ", tab.Title))
		}
	}
	b.WriteString("\n\n")
	selectedItem := m.tabContent[m.activeTab].SelectedItem()
	var i backend.Pkg
	if selectedItem != nil {
		p, ok := selectedItem.(backend.Pkg)
		if ok {
			i = p
		}
	}
	switch m.state {
	case stateInstalled:
		return fmt.Sprintf("%s was successfully installed", i.Name)
	case stateRemoved:
		return fmt.Sprintf("%s has been uninstalled", i.Name)
	case stateUpdated:
		return "Packages have been updated!"
	}

	b.WriteString(m.tabContent[m.activeTab].View())
	return b.String()
}
