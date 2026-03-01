package packagebrowser

import (
	"fmt"
	"strings"

	"github.com/ElitistNoob/pacdude/internal/tui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m *PackageBrowserModel) View() string {
	var b strings.Builder

	pkg := m.getSelectedPackage()
	switch m.state {
	case stateInstalled:
		return fmt.Sprintf("%s was successfully installed\n\n[space] continue [q] quit", pkg)
	case stateRemoved:
		return fmt.Sprintf("%s has been uninstalled", pkg)
	case stateUpdated:
		return "Packages have been updated!"
	}

	contentWidth := m.getContentWidth(styles.ListStyle)
	b.WriteString(styles.ListStyle.
		Width(contentWidth).
		Render(m.tabs.Active().View()))
	return m.RenderTabs() + b.String()
}

func (m *PackageBrowserModel) RenderTabs() string {
	var parts []string

	parts = append(parts, "╭─")
	for i, tab := range m.tabs.Tabs {
		if i == int(m.tabs.Index) {
			parts = append(parts, " "+styles.TabActive.Render(tab)+" ─")
		} else {
			parts = append(parts, " "+styles.TabInactive.Render(tab)+" ─")
		}
	}
	tabRow := lipgloss.JoinHorizontal(lipgloss.Top, parts...)

	contentWidth := m.getContentWidth(styles.ListStyle)
	rowWidth := lipgloss.Width(tabRow)

	line := lipgloss.NewStyle().
		Width(contentWidth).
		Render(strings.Repeat(
			"─",
			contentWidth-rowWidth) + "─╮",
		)

	if rowWidth >= contentWidth {
		return tabRow
	}

	return tabRow + line
}

func (m *PackageBrowserModel) getContentWidth(style lipgloss.Style) int {
	w, _ := style.GetFrameSize()
	return m.width - w
}
