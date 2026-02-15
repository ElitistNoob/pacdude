package packagebrowser

import (
	"strings"

	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *PackageBrowserModel) Update(msg tea.Msg) (app.Screen, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	// Window Resize Messages
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-v, msg.Height-h)
		m.state = stateReady

	// Keypress Messages
	case tea.KeyMsg:
		switch msg.String() {
		case "i":
			m.state = stateConfirm
			return m, nil
		case "y":
			m.state = stateInstall
			selectedPkg := m.list.SelectedItem()
			if selectedPkg != nil {
				p, ok := selectedPkg.(pkg)
				if ok {
					return m, m.Backend.Install(strings.Split(p.title, " ")[0])
				}
			}
		case "n", "q":
			m.state = stateReady
			return m, m.Backend.ListInstalled()
		}

	case backend.ResultMsg:
		output := m.parseOutput(msg.Output)
		items := make([]list.Item, len(output))
		for i, v := range output {
			items[i] = v
		}
		return m, m.list.SetItems(items)

	case backend.InstallResultMsg:
		if msg.Result.Err.Err != nil {
			m.state = stateError
			m.error = msg.Result.Err.Err.Error()
			return m, nil
		}
		m.state = stateComplete
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
