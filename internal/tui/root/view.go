package root

import (
	"fmt"
)

func (m model) View() string {
	// The header
	s := "What do you want to do?\n\n"

	// Iterate over our choices
	for i, cmd := range m.cmds {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s]\n", cursor, cmd.name)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
