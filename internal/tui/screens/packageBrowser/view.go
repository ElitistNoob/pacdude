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

		headerTabs := m.RenderTabs(m.managerTab, contentW)
		headerContent := styles.ContentStyle.
			Width(contentW).
			Padding(0, 1).
			Render(m.managerTab.Active().(*panels.TextPanel).Text)
		headerBlock := lipgloss.JoinVertical(lipgloss.Top, headerTabs, headerContent)

		bodyHeight := contentH - (lipgloss.Height(headerBlock) + 2)
		infoWidth := contentW / 4
		categoriesWidth := contentW - infoWidth

		m.tabs.Active().SetSize(categoriesWidth, bodyHeight)

		categoriesTabs := m.RenderTabs(m.tabs, categoriesWidth)
		categoriesContent := styles.ContentStyle.
			Width(categoriesWidth).
			Height(bodyHeight).
			Render(m.tabs.Active().View())
		categoriesBlock := lipgloss.JoinVertical(lipgloss.Top, categoriesTabs, categoriesContent)

		infoContent := lipgloss.NewStyle().
			Height(contentH-lipgloss.Height(headerBlock)-1).
			Width(infoWidth).
			Padding(1, 1).
			Border(lipgloss.RoundedBorder()).
			Render(m.infoPanel.Text)
		bodyBlock := lipgloss.JoinHorizontal(lipgloss.Bottom, categoriesBlock, infoContent)

		return lipgloss.JoinVertical(lipgloss.Top,
			headerBlock,
			bodyBlock,
		)
	}

	return ""
}

func (m *PackageBrowserModel) RenderTabs(tm *tabs.TabsModel, width int) string {
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

	rowWidth := lipgloss.Width(tabRow)
	remaining := max(width-rowWidth, 0)

	return tabRow + strings.Repeat("─", remaining) + "─╮"
}
