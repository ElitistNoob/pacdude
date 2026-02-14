package tui

import "fmt"

func (m model) View() string {
	switch m.state {
	case stateConfirmInstall:
		return fmt.Sprintf("Install %s?\n\n[y] Yes\n\n[n] No", m.selectedPkg)
	case stateInstalling:
		return fmt.Sprintf("Installing %s...", m.selectedPkg)
	case stateDone:
		return fmt.Sprintf("%s was successfully installed", m.selectedPkg)
	}
	return m.current.View()
}
