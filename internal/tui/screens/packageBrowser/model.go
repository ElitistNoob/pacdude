package packagebrowser

import (
	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	t "github.com/ElitistNoob/pacdude/internal/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	stateLoading state = iota
	stateReady
	stateInstalled
	stateRemoved
	stateUpdated
)

type PackageBrowserModel struct {
	backend backend.BackendInterface
	state   state
	tabs    *t.TabsModel
	error   string
	width   int
	height  int
}

func NewModel(b backend.BackendInterface) app.Screen {
	t := t.NewTabsModel()

	return &PackageBrowserModel{
		backend: b,
		tabs:    t,
	}
}

func (m *PackageBrowserModel) Init() tea.Cmd {
	return tea.Batch(m.tabs.Active().ToggleSpinner(), m.backend.ListInstalled())
}
