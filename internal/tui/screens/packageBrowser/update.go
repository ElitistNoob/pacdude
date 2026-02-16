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
		m.list.SetSize(msg.Width-v, msg.Height-h)
		m.state = stateReady

	// Keypress Messages
	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			switch msg.String() {
			case "enter":
				query := m.list.FilterValue()

				m.list.FilterInput.Blur()
				m.list.ResetFilter()
				m.list.SetFilterState(list.Unfiltered)
				m.list.SetShowFilter(false)
				m.list.SetShowTitle(true)

				m.list.Title = "Search Results: " + query

				return m, tea.Batch(m.list.ToggleSpinner(), m.Backend.Search(query))
			}
			break
		}
		switch {
		case key.Matches(msg, m.keys.install):
			selectedPkg := m.list.SelectedItem()
			if selectedPkg != nil {
				p, ok := selectedPkg.(backend.Pkg)
				if ok {
					return m, m.Backend.Install(strings.Split(p.Name, " ")[0])
				}
			}
		case key.Matches(msg, m.keys.remove):
			selectedPkg := m.list.SelectedItem()
			if selectedPkg != nil {
				p, ok := selectedPkg.(backend.Pkg)
				if ok {
					return m, m.Backend.Remove(strings.Split(p.Name, " ")[0])
				}
			}
		case key.Matches(msg, m.keys.updatable):
			return m, tea.Batch(m.list.ToggleSpinner(), m.Backend.ListUpgradable())
		case key.Matches(msg, m.keys.updateAll):
			return m, m.Backend.UpdateAll()
		case key.Matches(msg, m.keys.InstalledPackage):
			return m, tea.Batch(m.list.ToggleSpinner(), m.Backend.ListInstalled())
		}

	// Backend Messages
	case backend.ListInstalledPackagesMsg:
		m.state = stateReady
		m.list.StopSpinner()
		m.list.Title = "Installed Packages"
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
		m.list.Title = "Available Updates"
		m.list.StopSpinner()
		return m, m.setListItems(msg.Output)
	case backend.SearchPacmanPackagesMsg:
		m.list.StopSpinner()
		m.list.FilterInput.Focus()
		return m, m.setListItems(msg.Output)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
