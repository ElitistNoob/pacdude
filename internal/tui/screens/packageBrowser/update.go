package packagebrowser

import (
	"log"

	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *PackageBrowserModel) Update(msg tea.Msg) (app.Screen, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	// WINDOW RESIZE
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		h, v := docStyle.GetFrameSize()
		for i := range m.tabContent {
			m.tabContent[i].SetSize(msg.Width-v, msg.Height-h)
		}
		m.state = stateReady

	// NORMAL KEYS //
	case tea.KeyMsg:
		if m.tabContent[m.activeTab].FilterState() == list.Filtering {
			switch msg.String() {
			case "enter":
				query := m.tabContent[m.activeTab].FilterValue()

				m.tabContent[m.activeTab].FilterInput.Blur()
				m.tabContent[m.activeTab].ResetFilter()
				m.tabContent[m.activeTab].SetFilterState(list.Unfiltered)
				m.tabContent[m.activeTab].SetShowFilter(true)
				m.tabContent[m.activeTab].SetShowTitle(true)

				m.tabContent[m.activeTab].Title = "Search Results: " + query

				return m, tea.Batch(m.tabContent[m.activeTab].ToggleSpinner(), m.Backend.Search(query))
			}
			break
		}

		// LIST KEYS //
		switch {
		case key.Matches(msg, m.keys.installedPackage):
			m.activeTab = 0
			return m, m.loadPackageData(m.Backend.ListInstalled())
		case key.Matches(msg, m.keys.install):
			pkg := m.getSelectedPackage()
			return m, m.Backend.Install(pkg)
		case key.Matches(msg, m.keys.updatable):
			m.activeTab = 1
			return m, m.loadPackageData(m.Backend.ListUpgradable())
		case key.Matches(msg, m.keys.updateAll):
			return m, m.Backend.UpdateAll()
		case key.Matches(msg, m.keys.uninstall):
			pkg := m.getSelectedPackage()
			return m, m.Backend.Remove(pkg)
		}

	// BACKEND MESSAGES //
	case backend.ListInstalledPackagesMsg:
		m.state = stateReady
		m.tabContent[m.activeTab].StopSpinner()
		return m, m.setListItems(msg.Output)
	case backend.InstallPackageResultMsg:
		if msg.Err.Err != nil {
			log.Printf("installation error: %v", msg.Err.Err)
			m.error = msg.Err.Err.Error()
			return m, nil
		}
		m.state = stateInstalled
		return m, nil
	case backend.ListAvailableUpdatesMsg:
		m.tabContent[m.activeTab].StopSpinner()
		return m, m.setListItems(msg.Output)
	case backend.RemovePackageResultMsg:
		if msg.Err.Err != nil {
			log.Printf("error uninstalling package: %v", msg.Err.Err)
			m.error = msg.Err.Err.Error()
			return m, nil
		}
		m.state = stateRemoved
		return m, nil
	case backend.UpdateAllMsg:
		if msg.Err.Err != nil {
			log.Printf("error updating packages: %v", msg.Err.Err)
			m.error = msg.Err.Err.Error()
			return m, nil
		}
		m.state = stateUpdated
		return m, nil
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
