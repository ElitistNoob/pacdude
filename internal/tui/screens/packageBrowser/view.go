package packagebrowser

import (
	"fmt"
)

func (m *PackageBrowserModel) View() string {
	selectedPkg := m.list.SelectedItem()
	switch m.state {
	case stateLoading:
		return "\n Initializing..."
	case stateConfirm:
		return fmt.Sprintf("Install %s\n\n[y] Yes\n\n[n] No", selectedPkg)
	case stateInstall:
		return fmt.Sprintf("Installing %s...", selectedPkg)
	case stateComplete:
		return fmt.Sprintf("%s was installed successfully\n\nPress [q] to continue", selectedPkg)
	case stateError:
		return fmt.Sprintf(
			"An error occured while trying to install %s\n\nErr: %s\n\nPress [q] to continue",
			selectedPkg,
			m.error,
		)
	}
	return m.list.View()
}
