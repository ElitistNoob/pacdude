package packagebrowser

import (
	"strings"

	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m *PackageBrowserModel) Update(msg tea.Msg) (app.Screen, tea.Cmd) {
	if m.showModal {
		switch msg := msg.(type) {
		case backend.ResultMsg:
			m.choices = parseOutput(msg.Output)
			m.viewport.SetContent(m.renderContent())

		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.queryResult = m.textInput.Value()
				m.showModal = false
				m.textInput.Blur()
				return m, m.Backend.Search(m.queryResult)
			case "esc":
				m.showModal = false
				m.textInput.Blur()
				return m, nil
			}
		}

		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	var cmds []tea.Cmd
	switch msg := msg.(type) {

	// Window Resize Messages
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		switch m.state {
		case stateLoading:
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.state = stateReady
		case stateReady:
			if m.viewport.Width == 0 {
				m.viewport.Width = msg.Width
				m.viewport.Height = msg.Height - verticalMarginHeight
				m.viewport.SetContent(m.renderContent())
			}
		}

	// Keypress Messages
	case tea.KeyMsg:
		if len(m.choices) == 0 {
			return m, nil
		}
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "/":
			m.showModal = true
			m.textInput.Focus()
			m.textInput.SetValue("")
			return m, textinput.Blink
		case "i":
			m.state = stateConfirm
			return m, nil
		case "y":
			m.state = stateInstall
			return m, m.Backend.Install(m.selection)
		case "n", "q":
			m.state = stateReady
			return m, m.Backend.ListInstalled()
		}
		if m.cursor > len(m.choices) {
			m.cursor = 0
		}
		if len(m.choices) > 0 {
			m.selection = strings.Split(m.choices[m.cursor].title, " ")[0]
			m.viewport.SetContent(m.renderContent())
			m.syncViewportScroll()
		}

	case backend.ResultMsg:
		m.choices = parseOutput(msg.Output)
		m.viewport.SetContent(m.renderContent())
	case backend.InstallResultMsg:
		if msg.Result.Err.Err != nil {
			m.state = stateError
			m.error = msg.Result.Err.Err.Error()
			return m, nil
		}
		m.state = stateComplete
		// m.lastOutput = outputMsg(msg.Output)
		return m, nil
	}

	var vpCmd tea.Cmd
	m.viewport, vpCmd = m.viewport.Update(msg)
	cmds = append(cmds, vpCmd)
	return m, tea.Batch(cmds...)
}
