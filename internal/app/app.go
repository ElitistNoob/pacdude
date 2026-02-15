package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	Current Screen
	width   int
	height  int
}

func (m AppModel) Init() tea.Cmd {
	return m.Current.Init()
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	case ChangeScreenMsg:
		m.Current = msg.NewScreen
		initCmd := m.Current.Init()

		_, sizeCmd := m.Current.Update(tea.WindowSizeMsg{
			Width:  m.width,
			Height: m.height,
		})
		return m, tea.Batch(initCmd, sizeCmd)
	}

	var cmd tea.Cmd
	m.Current, cmd = m.Current.Update(msg)
	return m, cmd
}

func (m AppModel) View() string {
	return m.Current.View()
}
