package packagebrowser

import (
	"fmt"
	"strings"

	"github.com/ElitistNoob/pacdude/internal/tui/styles"
)

func (m *PackageBrowserModel) View() string {
	var b strings.Builder
	for i := range m.tabs.Tabs {
		style := styles.TabInactive
		if i == int(m.tabs.Index) {
			style = styles.TabActive
		}

		b.WriteString(style.Render(m.tabs.Tabs[i].Title))
		b.WriteString(" ")
	}
	b.WriteString("\n\n")

	pkg := m.getSelectedPackage()
	switch m.state {
	case stateInstalled:
		return fmt.Sprintf("%s was successfully installed\n\n[space] continue [q] quit", pkg)
	case stateRemoved:
		return fmt.Sprintf("%s has been uninstalled", pkg)
	case stateUpdated:
		return "Packages have been updated!"
	}

	b.WriteString(m.tabs.Active().View())
	return b.String()
}
