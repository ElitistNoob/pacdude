package packagebrowser

import (
	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func (i pkg) Title() string       { return i.title }
func (i pkg) Description() string { return i.desc }
func (i pkg) FilterValue() string { return i.title }

type PackageBrowserModel struct {
	Backend backend.BackendInterface
	state   state
	list    list.Model
	error   string
	width   int
	height  int
}

func NewModel(b backend.BackendInterface) app.Screen {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Packges"

	return &PackageBrowserModel{
		Backend: b,
		list:    l,
	}
}

func (m *PackageBrowserModel) Init() tea.Cmd {
	return m.Backend.ListInstalled()
}
