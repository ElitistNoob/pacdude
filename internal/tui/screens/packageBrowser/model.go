package packagebrowser

import (
	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	stateLoading state = iota
	stateReady
	stateConfirm
	stateInstall
	stateComplete
	stateError
)

type pkg struct {
	title, desc string
}

type PackageBrowserModel struct {
	Backend     backend.BackendInterface
	state       state
	choices     []pkg
	output      string
	cursor      int
	selection   string
	showModal   bool
	textInput   textinput.Model
	queryResult string
	viewport    viewport.Model
	error       string
	width       int
	height      int
}

func NewModel(b backend.BackendInterface) app.Screen {
	ti := textinput.New()
	ti.Placeholder = "Search Package"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40

	return &PackageBrowserModel{
		Backend:   b,
		choices:   []pkg{},
		textInput: ti,
		showModal: false,
	}
}

func (m *PackageBrowserModel) Init() tea.Cmd {
	return m.Backend.ListInstalled()
}

func (m *PackageBrowserModel) syncViewportScroll() {
	lineHeight := 3
	topOfItem := m.cursor * lineHeight
	bottomOfItem := topOfItem + (lineHeight - 1)

	top := m.viewport.YOffset
	bottom := m.viewport.YOffset + m.viewport.Height - 1

	if topOfItem < top {
		m.viewport.SetYOffset(topOfItem)
	} else if bottomOfItem > bottom {
		m.viewport.SetYOffset(bottomOfItem - m.viewport.Height + 1)
	}
}
