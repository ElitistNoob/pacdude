package packagebrowser

import (
	"strings"

	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *PackageBrowserModel) Update(msg tea.Msg) (app.Screen, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	// Window Resize Messages
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		h, v := docStyle.GetFrameSize()
		m.tabContent[m.activeTab].SetSize(msg.Width-v, msg.Height-h)
		m.state = stateReady

	// Keypress Messages
	case tea.KeyMsg:
		if m.tabContent[m.activeTab].FilterState() == list.Filtering {
			switch msg.String() {
			case "enter":
				query := m.tabContent[m.activeTab].FilterValue()

				m.tabContent[m.activeTab].FilterInput.Blur()
				m.tabContent[m.activeTab].ResetFilter()
				m.tabContent[m.activeTab].SetFilterState(list.Unfiltered)
				m.tabContent[m.activeTab].SetShowFilter(false)
				m.tabContent[m.activeTab].SetShowTitle(true)

				// m..Title = "Search Results: " + query

				return m, tea.Batch(m.tabContent[m.activeTab].ToggleSpinner(), m.Backend.Search(query))
			}
			break
		}
		switch {
		case key.Matches(msg, m.keys.install):
			m.activeTab = 0
			// selectedPkg := m.tabContent[m.activeTab].SelectedItem()
			// if selectedPkg != nil {
			// 	p, ok := selectedPkg.(backend.Pkg)
			// 	if ok {
			// 		return m, m.Backend.Install(strings.Split(p.Name, " ")[0])
			// 	}
			// }
		case key.Matches(msg, m.keys.remove):
			selectedPkg := m.tabContent[m.activeTab].SelectedItem()
			if selectedPkg != nil {
				p, ok := selectedPkg.(backend.Pkg)
				if ok {
					return m, m.Backend.Remove(strings.Split(p.Name, " ")[0])
				}
			}
		case key.Matches(msg, m.keys.updatable):
			m.activeTab = 1
			return m, tea.Batch(func() tea.Msg { return m.tabContent[m.activeTab] }, m.tabContent[m.activeTab].ToggleSpinner(), m.Backend.ListUpgradable())
		case key.Matches(msg, m.keys.updateAll):
			return m, m.Backend.UpdateAll()
		case key.Matches(msg, m.keys.InstalledPackage):
			return m, tea.Batch(m.tabContent[m.activeTab].ToggleSpinner(), m.Backend.ListInstalled())
		}

	// Backend Messages
	case backend.ListInstalledPackagesMsg:
		m.state = stateReady
		m.tabContent[m.activeTab].StopSpinner()
		// m.list.Title = "Installed Packages"
		return m, m.setListItems(msg.Output)
	case backend.InstallPackageResultMsg:
		if msg.Err.Err != nil {
			m.error = msg.Err.Err.Error()
			return m, nil
		}

		m.state = stateInstalled
		return m, nil
	case backend.RemovePackageResultMsg:
		if msg.Err.Err != nil {
			m.error = msg.Err.Err.Error()
			return m, nil
		}

		m.state = stateRemoved
		return m, nil

	case backend.UpdateAllMsg:
		if msg.Err.Err != nil {
			m.error = msg.Err.Err.Error()
			return m, nil
		}
		m.state = stateUpdated
		return m, nil
	case backend.ListAvailableUpdatesMsg:
		// m.list.Title = "Available Updates"
		m.tabContent[m.activeTab].StopSpinner()
		return m, m.setListItems(msg.Output)
	case backend.SearchPacmanPackagesMsg:
		m.tabContent[m.activeTab].StopSpinner()
		m.tabContent[m.activeTab].FilterInput.Focus()
		return m, m.setListItems(msg.Output)
	}

	var cmd tea.Cmd
	m.tabContent[m.activeTab], cmd = m.tabContent[m.activeTab].Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
