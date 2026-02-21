package packagebrowser

import (
	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	stateLoading state = iota
	stateReady
	stateInstalled
	stateRemoved
	stateUpdated
)

type tab int

const (
	installed tab = iota
	updatable
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type listKeyMap struct {
	install          key.Binding
	uninstall        key.Binding
	updatable        key.Binding
	updateAll        key.Binding
	installedPackage key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		installedPackage: key.NewBinding(
			key.WithKeys("I"),
			key.WithHelp("I", "show installed packages"),
		),
		install: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "install"),
		),
		updatable: key.NewBinding(
			key.WithKeys("U"),
			key.WithHelp("U", "show available updates"),
		),
		uninstall: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "uninstall package"),
		),
	}
}

type PackageBrowserModel struct {
	Backend    backend.BackendInterface
	state      state
	tabs       []string
	activeTab  tab
	tabContent []list.Model
	keys       *listKeyMap
	error      string
	width      int
	height     int
}

func NewModel(b backend.BackendInterface) app.Screen {
	tabs := []string{"Installed (I)", "Available Updates (U)"}
	tabContent := make([]list.Model, len(tabs))
	listKey := newListKeyMap()
	for i, tab := range tabs {
		l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
		l.Title = tab
		l.SetShowTitle(false)
		l.SetShowStatusBar(true)
		l.AdditionalFullHelpKeys = func() []key.Binding {
			return []key.Binding{
				listKey.installedPackage,
				listKey.install,
				listKey.updatable,
				listKey.updateAll,
				listKey.uninstall,
			}
		}

		tabContent[i] = l
	}

	return &PackageBrowserModel{
		Backend:    b,
		tabs:       tabs,
		tabContent: tabContent,
		keys:       listKey,
		activeTab:  0,
	}
}

func (m *PackageBrowserModel) Init() tea.Cmd {
	return tea.Batch(m.tabContent[m.activeTab].ToggleSpinner(), m.Backend.ListInstalled())
}
