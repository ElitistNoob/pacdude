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
	list    list.Model
	keys    *listKeyMap
	error   string
	width   int
	height  int
}

func NewModel(b backend.BackendInterface) app.Screen {
	listKey := newListKeyMap()
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Installed Packages"
	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKey.install,
			listKey.remove,
			listKey.updatable,
			listKey.updateAll,
			listKey.InstalledPackage,
		}
	}

	return &PackageBrowserModel{
		Backend: b,
		list:    l,
		keys:    listKey,
	}
}

func (m *PackageBrowserModel) Init() tea.Cmd {
	return m.Backend.ListInstalled()
}
