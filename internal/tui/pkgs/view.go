package pkgs

import (
	"fmt"
	"strings"
)

func (m model) View() string {
	return renderView(m)
}

func renderView(m model) string {
	var str strings.Builder
	lines := make([]string, 0, len(m.choices))
	for i, pkg := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		lines = append(lines, fmt.Sprintf("[%s] %s\n%s\n", cursor, pkg.title, pkg.desc))
	}

	lines = append(lines, "Press q to quit\n")

	str.WriteString(strings.Join(lines, "\n"))
	return str.String()
}
