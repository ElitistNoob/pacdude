package pkgs

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	choices []pkg
	cursor  int
}

type pkg struct {
	title, desc string
}

func NewPkgsModel() *model {
	return &model{
		choices: []pkg{},
		cursor:  0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) SetPackages(output []byte) {
	m.choices = parseOutput(output)
}
