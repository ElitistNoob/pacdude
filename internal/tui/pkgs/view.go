package pkgs

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()
)

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s",
		m.headerView(),
		m.viewport.View(),
		m.footerView(),
	)
	// return fmt.Sprintf("Cursor: %d | YOffset: %d\n%s", m.cursor, m.viewport.YOffset, m.viewport.View())
}

func (m model) renderContent() string {
	lines := make([]string, 0, len(m.choices))
	for i, pkg := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		lines = append(lines, fmt.Sprintf("[%s] %s\n%s\n", cursor, pkg.title, pkg.desc))
	}

	return strings.Join(lines, "\n")
}

func (m model) headerView() string {
	title := titleStyle.Render("Packages")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
