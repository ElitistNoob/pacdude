package packagebrowser

import (
	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	panels "github.com/ElitistNoob/pacdude/internal/tui/components"
	"github.com/ElitistNoob/pacdude/internal/tui/components/tabs"
	t "github.com/ElitistNoob/pacdude/internal/tui/components/tabs"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	state int
)

const (
	stateLoading state = iota
	stateReady
	stateInstalled
	stateRemoved
	stateUpdated
	stateError
)

type Panel interface {
	Update(tea.Msg) (Panel, tea.Cmd)
	View() string
	SetSize(w, h int)
}

type PackageBrowserModel struct {
	backend    backend.BackendInterface
	state      state
	tabs       *t.TabsModel
	managerTab *t.TabsModel
	width      int
	height     int
	keys       browserKeyMap
}

type browserKeyMap struct {
	viewAll              key.Binding
	viewInstalled        key.Binding
	viewAvailableUpdates key.Binding
	installPackage       key.Binding
	removePackage        key.Binding
	updateAllPackages    key.Binding
	nextTab              key.Binding
	prevTab              key.Binding
}

func newBrowserKeyMap() *browserKeyMap {
	return &browserKeyMap{
		viewAll: key.NewBinding(
			key.WithKeys("A"),
			key.WithHelp("A", "All Packages"),
		),
		viewInstalled: key.NewBinding(
			key.WithKeys("I"),
			key.WithHelp("I", "Installed Packages"),
		),
		viewAvailableUpdates: key.NewBinding(
			key.WithKeys("U"),
			key.WithHelp("U", "Available Updates"),
		),
		installPackage: key.NewBinding(
			key.WithKeys("i"),
			key.WithHelp("i", "Install"),
		),
		removePackage: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "Uninstall"),
		),
		updateAllPackages: key.NewBinding(
			key.WithKeys("u"),
			key.WithHelp("u", "Uninstall"),
		),
		nextTab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("Tab", "Next Tab"),
		),
		prevTab: key.NewBinding(
			key.WithKeys("shift+tab"),
			key.WithHelp("Shift + Tab", "Prev Tab"),
		),
	}
}

func NewModel(b backend.BackendInterface, index int) app.Screen {
	pacmanPanel := &panels.TextPanel{Text: "Pacman - I use Arch btw"}
	flatpakPanel := &panels.TextPanel{Text: "Flatpak"}
	brewPanel := &panels.TextPanel{Text: "Brew"}

	managersTabs := tabs.NewTabsModel(
		[]string{"Pacman", "Flatpak", "Brew"},
		[]tabs.TabContent{pacmanPanel, flatpakPanel, brewPanel},
	)
	managersTabs.Index = index

	allList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	installedList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	updatesList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)

	allPanel := panels.NewListPanel(allList)
	installedPanel := panels.NewListPanel(installedList)
	updatesPanel := panels.NewListPanel(updatesList)
	t := t.NewTabsModel(
		[]string{"All (A)", "Installed (I)", "Available Updates (U)"},
		[]t.TabContent{
			allPanel,
			installedPanel,
			updatesPanel,
		},
	)

	return &PackageBrowserModel{
		backend:    b,
		managerTab: managersTabs,
		tabs:       t,
		keys:       *newBrowserKeyMap(),
	}
}

func (m *PackageBrowserModel) Init() tea.Cmd {
	items := m.backend.ListAll
	return tea.Batch(runBackend(items))
}
