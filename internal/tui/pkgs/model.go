package pkgs

import (
	msg "github.com/ElitistNoob/pacdude/internal/tui/messages"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	ready     bool
	choices   []pkg
	cursor    int
	selection string
	showModal bool
	textInput textinput.Model
	result    string
	viewport  viewport.Model
}

type pkg struct {
	title, desc string
}

func NewPkgsModel() *model {
	ti := textinput.New()
	ti.Placeholder = "Search Package"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40
	return &model{
		choices:   []pkg{},
		cursor:    0,
		viewport:  viewport.Model{},
		textInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) SetPackages(output []byte) {
	m.choices = parseOutput(output)
}

func (m *model) syncViewportScroll() {
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

func (m *model) handleInstallPkgMsg(args []string) tea.Cmd {
	return func() tea.Msg {
		return msg.InstallPkgMsg{Args: args}
	}
}

func (m *model) handleExecDoneMsg() tea.Cmd {
	return func() tea.Msg {
		return msg.ExecDoneMsg{}
	}
}
