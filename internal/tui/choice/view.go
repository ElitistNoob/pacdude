package choice

import (
	"fmt"
	"strings"
)

func (m model) View() string {
	return renderView(m)

}

func renderView(m model) string {
	var menu strings.Builder
	lines := make([]string, 0, len(m.choices)-1)
	lines = append(lines, "What do you want to do?\n\n")
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		lines = append(lines, fmt.Sprintf("[%s] %s\n", cursor, choice))
	}

	lines = append(lines, "Press q to quit\n")

	menu.WriteString(strings.Join(lines, "\n"))
	return menu.String()
}
