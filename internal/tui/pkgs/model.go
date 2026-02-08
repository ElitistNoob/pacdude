package pkgs

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices   []pkg
	cursor    int
	viewport  viewport.Model
	ready     bool
	textInput textinput.Model
	showModal bool
	result    string
	width     int
	height    int
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
