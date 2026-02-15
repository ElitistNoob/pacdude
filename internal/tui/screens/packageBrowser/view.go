package packagebrowser

import (
	"fmt"

	"github.com/ElitistNoob/pacdude/internal/backend"
)

func (m *PackageBrowserModel) View() string {
	selectedItem := m.list.SelectedItem()
	var i backend.Pkg
	if selectedItem != nil {
		p, ok := selectedItem.(backend.Pkg)
		if ok {
			i = p
		}
	}
	switch m.state {
	case stateInstalled:
		return fmt.Sprintf("%s was successfully installed", i.Name)
	case stateRemoved:
		return fmt.Sprintf("%s has been uninstalled", i.Name)
	case stateUpdated:
		return "Packages have been updated!"
	}
	return m.list.View()
}
