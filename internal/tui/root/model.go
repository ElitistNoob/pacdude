package root

import (
	"github.com/ElitistNoob/pacdude/internal/app"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cmds   []cmd
	cursor int
	output outputMsg
	err    errMsg
}

type outputMsg string
type errMsg struct {
	err error
}

type cmd struct {
	name string
	run  func() (string, error)
}

func InitialModel() model {
	return model{
		cmds: []cmd{
			{
				name: "Installed Packages",
				run:  app.InstalledPackages,
			},
		},
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
