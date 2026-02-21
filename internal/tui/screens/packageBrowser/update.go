package packagebrowser

import (
	"log"

	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	tabs "github.com/ElitistNoob/pacdude/internal/tui/components"
	"github.com/ElitistNoob/pacdude/internal/tui/styles"
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
		h, v := styles.DocStyle.GetFrameSize()
		m.tabs.SetSize(msg.Width-v, msg.Height-h)
		m.state = stateReady

	// NORMAL KEYS //
	case tea.KeyMsg:
		if m.tabs.Active().FilterState() == list.Filtering {
			switch msg.String() {
			case "enter":
				query := m.tabs.Active().FilterValue()

				m.tabs.Active().FilterInput.Blur()
				m.tabs.Active().ResetFilter()
				m.tabs.Active().SetFilterState(list.Unfiltered)
				m.tabs.Active().SetShowFilter(true)
				m.tabs.Active().SetShowTitle(true)

				m.tabs.Active().Title = "Search Results: " + query

				return m, tea.Batch(m.tabs.Active().ToggleSpinner(), m.backend.Search(query))
			}
			break
		}

		// LIST KEYS //
		switch {
		case key.Matches(msg, m.tabs.Keys.InstalledPackage):
			m.tabs.Index = tabs.Installed
			return m, m.loadPackageData(m.backend.ListInstalled())
		case key.Matches(msg, m.tabs.Keys.Install):
			pkg := m.getSelectedPackage()
			return m, m.backend.Install(pkg)
		case key.Matches(msg, m.tabs.Keys.Updatable):
			m.tabs.Index = tabs.Updatable
			return m, m.loadPackageData(m.backend.ListUpgradable())
		case key.Matches(msg, m.tabs.Keys.UpdateAll):
			return m, m.backend.UpdateAll()
		case key.Matches(msg, m.tabs.Keys.Uninstall):
			pkg := m.getSelectedPackage()
			return m, m.backend.Remove(pkg)
		}

	// BACKEND MESSAGES //
	case backend.ListInstalledPackagesMsg:
		m.state = stateReady
		m.tabs.Active().StopSpinner()
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
		m.tabs.Active().StopSpinner()
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
		m.tabs.Active().StopSpinner()
		m.tabs.Active().FilterInput.Focus()
		return m, m.setListItems(msg.Output)
	}

	var cmd tea.Cmd
	*m.tabs.Active(), cmd = m.tabs.Active().Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
