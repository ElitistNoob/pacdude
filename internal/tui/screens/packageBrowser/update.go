package packagebrowser

import (
	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	tabs "github.com/ElitistNoob/pacdude/internal/tui/components"
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
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
				query := m.tabs.Query()

				m.tabs.Active().FilterInput.Blur()
				m.tabs.Active().ResetFilter()
				m.tabs.Active().SetFilterState(list.Unfiltered)
				m.tabs.Active().SetShowTitle(true)
				m.tabs.Active().SetShowFilter(true)
				m.tabs.Active().Title = "Search Results: " + query

				fn := runBackend(func() backend.ResultMsg {
					return m.backend.Search(query)
				})

				return m, tea.Batch(m.tabs.Active().ToggleSpinner(), fn)
			}
			break
		}

		// LIST KEYS //
		switch {
		case key.Matches(msg, m.tabs.Keys.InstalledPackage):
			m.tabs.Index = tabs.Installed
			return m, runBackend(m.backend.ListInstalled)
		case key.Matches(msg, m.tabs.Keys.Install):
			pkg := m.getSelectedPackage()
			return m, runBackend(func() backend.ResultMsg {
				return m.backend.Install(pkg)
			})
		case key.Matches(msg, m.tabs.Keys.Updatable):
			m.tabs.Index = tabs.Updatable
			return m, runBackend(m.backend.ListUpgradable)
		case key.Matches(msg, m.tabs.Keys.UpdateAll):
			return m, runBackend(m.backend.UpdateAll)
		case key.Matches(msg, m.tabs.Keys.Uninstall):
			pkg := m.getSelectedPackage()
			return m, runBackend(func() backend.ResultMsg {
				return m.backend.Remove(pkg)
			})
		}

		// BACKEND MESSAGES //
	case messages.ActionMsg:
		switch msg.Type {
		case messages.ActionInstalledLoaded:
			m.state = stateReady
			m.tabs.Active().StopSpinner()
			pkgs := m.backend.ParseOutput(msg.Payload.(backend.OutputMsg))
			return m, m.setListItems(pkgs)
		case messages.ActionUpdatesLoaded:
			m.tabs.Active().StopSpinner()
			pkgs := m.backend.ParseOutput(msg.Payload.(backend.OutputMsg))
			return m, m.setListItems(pkgs)
		case messages.ActionPackageInstalled:
			m.state = stateInstalled
			return m, nil
		case messages.ActionUpdatedAll:
			m.state = stateUpdated
			return m, nil
		case messages.ActionPackageRemoved:
			m.state = stateRemoved
			return m, nil
		case messages.ActionSearchLoaded:
			m.tabs.Active().StopSpinner()
			m.tabs.Active().FilterInput.Focus()
			pkgs := m.backend.ParseOutput(msg.Payload.(backend.OutputMsg))
			return m, m.setListItems(pkgs)
		case messages.ActionError:
			m.state = stateError
			return m, nil
		}
	}

	var cmd tea.Cmd
	*m.tabs.Active(), cmd = m.tabs.Active().Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
