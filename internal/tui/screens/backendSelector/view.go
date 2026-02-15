package backendselector

import (
	"fmt"
	"strings"
)

func (m *backendSelectorModel) View() string {
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

	menu.WriteString(strings.Join(lines, "\n"))
	return menu.String()
}
