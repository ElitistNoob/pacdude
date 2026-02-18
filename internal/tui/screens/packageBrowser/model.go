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

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type listKeyMap struct {
	install          key.Binding
	remove           key.Binding
	updatable        key.Binding
	updateAll        key.Binding
	InstalledPackage key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		install: key.NewBinding(
			key.WithKeys("i"),
			key.WithHelp("i", "install"),
		),
		InstalledPackage: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "install"),
		),
		remove: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "uninstall"),
		),
		updatable: key.NewBinding(
			key.WithKeys("u"),
			key.WithHelp("u", "show available updates"),
		),
		updateAll: key.NewBinding(
			key.WithKeys("U"),
			key.WithHelp("U", "update all packages"),
		),
	}
}

type PackageBrowserModel struct {
	Backend backend.BackendInterface
	state   state
	// list    list.Model
	tabContent []list.Model
	activeTab  int
	keys       *listKeyMap
	error      string
	width      int
	height     int
}

func NewModel(b backend.BackendInterface) app.Screen {
	tabs := []string{"Installed", "Available Updates"}
	tabContent := make([]list.Model, 0, len(tabs))
	listKey := newListKeyMap()
	for _, tab := range tabs {
		l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
		l.Title = tab
		l.SetShowTitle(false)
		l.SetShowStatusBar(true)
		l.AdditionalFullHelpKeys = func() []key.Binding {
			return []key.Binding{
				listKey.install,
				listKey.remove,
				listKey.updatable,
				listKey.updateAll,
				listKey.InstalledPackage,
			}
		}

		tabContent = append(tabContent, l)
	}

	return &PackageBrowserModel{
		Backend:    b,
		tabContent: tabContent,
		keys:       listKey,
		activeTab:  0,
	}
}

func (m *PackageBrowserModel) Init() tea.Cmd {
	return tea.Batch(m.tabContent[m.activeTab].ToggleSpinner(), m.Backend.ListInstalled())
}
