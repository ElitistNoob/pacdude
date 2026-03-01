package packagebrowser

import (
	"fmt"
	"strings"

	panels "github.com/ElitistNoob/pacdude/internal/tui/components"
	"github.com/ElitistNoob/pacdude/internal/tui/components/tabs"
	"github.com/ElitistNoob/pacdude/internal/tui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m *PackageBrowserModel) View() string {
	pkg := m.getSelectedPackage()
	switch m.state {
	case stateInstalled:
		return fmt.Sprintf("%s was successfully installed\n\n[space] continue [q] quit", pkg)
	case stateRemoved:
		return fmt.Sprintf("%s has been uninstalled", pkg)
	case stateUpdated:
		return "Packages have been updated!"
	}

	if m.state == stateReady {
		contentW, contentH := m.getContentSize(styles.ContentStyle)

		headerTabs := m.RenderTabs(m.managerTab)
		headerContent := styles.ContentStyle.
			Width(contentW).
			Padding(0, 1).
			Render(m.managerTab.Active().(*panels.TextPanel).Text)
		headerBlock := lipgloss.JoinVertical(lipgloss.Top, headerTabs, headerContent)

		bodyHeight := contentH - (lipgloss.Height(headerBlock) + 2)
		m.tabs.Active().SetSize(contentW, bodyHeight)

		bodyTabs := m.RenderTabs(m.tabs)
		bodyContent := styles.ContentStyle.
			Width(contentW).
			Height(bodyHeight).
			Render(m.tabs.Active().View())
		bodyBlock := lipgloss.JoinVertical(lipgloss.Top, bodyTabs, bodyContent)

		return lipgloss.JoinVertical(lipgloss.Top,
			headerBlock,
			bodyBlock,
		)
	}

	return ""
}

func (m *PackageBrowserModel) RenderTabs(tm *tabs.TabsModel) string {
	var parts []string

	parts = append(parts, "╭─")
	for i, tab := range tm.Tabs {
		if i == tm.Index {
			parts = append(parts, " "+styles.TabActive.Render(tab)+" ─")
		} else {
			parts = append(parts, " "+styles.TabInactive.Render(tab)+" ─")
		}
	}
	tabRow := lipgloss.JoinHorizontal(lipgloss.Top, parts...)

	contentWidth, _ := m.getContentSize(styles.ContentStyle)
	rowWidth := lipgloss.Width(tabRow)
	remaining := max(contentWidth-rowWidth, 0)

	return tabRow + strings.Repeat("─", remaining) + "─╮"
}
