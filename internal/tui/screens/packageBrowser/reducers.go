package packagebrowser

import (
	"github.com/ElitistNoob/pacdude/internal/backend"

	listpanel "github.com/ElitistNoob/pacdude/internal/tui/components"
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
	"github.com/ElitistNoob/pacdude/internal/tui/styles"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *PackageBrowserModel) reduceWindowSize(msg tea.WindowSizeMsg) tea.Cmd {
	m.width = msg.Width
	m.height = msg.Height

	h, v := styles.DocStyle.GetFrameSize()
	m.tabs.SetSize(msg.Width-v, msg.Height-h)

	m.state = stateReady
	return nil
}

func (m *PackageBrowserModel) reduceKeys(msg tea.KeyMsg) tea.Cmd {
	switch msg.String() {
	}
	switch {
	case key.Matches(msg, m.keys.viewAll):
		m.tabs.Index = 0
		return runBackend(m.backend.ListAll)

	case key.Matches(msg, m.keys.viewInstalled):
		m.tabs.Index = 1
		return runBackend(m.backend.ListInstalled)

	case key.Matches(msg, m.keys.viewAvailableUpdates):
		m.tabs.Index = 2
		return runBackend(m.backend.ListUpgradable)

	case key.Matches(msg, m.keys.installPackage):
		pkg := m.getSelectedPackage()
		return runBackend(func() backend.ResultMsg {
			return m.backend.Install(pkg)
		})

	case key.Matches(msg, m.keys.updateAllPackages):
		return runBackend(m.backend.UpdateAll)

	case key.Matches(msg, m.keys.removePackage):
		pkg := m.getSelectedPackage()
		return runBackend(func() backend.ResultMsg {
			return m.backend.Remove(pkg)
		})

	case key.Matches(msg, m.keys.nextTab):
		m.managerTab.NextTab()
		return m.setBackend(m.managerTab.Index)

	case key.Matches(msg, m.keys.prevTab):
		m.managerTab.PrevTab()
		return m.setBackend(m.managerTab.Index)
	}

	return nil
}

func (m *PackageBrowserModel) reduceActions(msg messages.ActionMsg) tea.Cmd {
	switch msg.Type {
	case messages.ActionPackagesLoaded:
		m.state = stateReady
		pkgs := m.backend.ParseOutput(msg.Payload.(backend.OutputMsg))
		return tea.Batch(m.setListItems(pkgs), m.loadActive())

	case messages.ActionPackageInstalled:
		m.state = stateInstalled

	case messages.ActionUpdatedAll:
		m.state = stateUpdated

	case messages.ActionPackageRemoved:
		m.state = stateRemoved

	case messages.ActionSearchLoaded:
		m.tabs.TabContent[m.tabs.Index].(*listpanel.ListPanel).List.FilterInput.Focus()
		pkgs := m.backend.ParseOutput(msg.Payload.(backend.OutputMsg))
		return m.setListItems(pkgs)

	case messages.ActionError:
		m.state = stateError
	}

	return nil
}
