package packagebrowser

import "fmt"

func (m *PackageBrowserModel) View() string {
	selectedItem := m.list.SelectedItem()
	var item pkg
	if selectedItem != nil {
		p, ok := selectedItem.(pkg)
		if ok {
			item = p
		}
	}
	switch m.state {
	case stateInstalled:
		return fmt.Sprintf("%s was successfully installed", item.title)
	case stateRemoved:
		return fmt.Sprintf("%s has been uninstalled", item.title)
	case stateUpdated:
		return "Packages have been updated!"
	}
	return m.list.View()
}
