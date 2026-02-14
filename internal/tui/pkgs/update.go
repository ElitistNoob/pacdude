package pkgs

import (
	"strings"

	hdlr "github.com/ElitistNoob/pacdude/internal/handlers"
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var args []string

	if m.showModal {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.result = m.textInput.Value()
				args = []string{"-Ss", m.result}
				m.showModal = false
				m.textInput.Blur()
				return m, hdlr.HandleShowPackageScreen(args)
			case "esc":
				m.showModal = false
				m.textInput.Blur()
				return m, nil
			}
		}

		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	switch msg := msg.(type) {

	// Window Resize Messages
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
			m.viewport.SetContent(m.renderContent())
		}

	// Keypress Messages
	case tea.KeyMsg:
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
			args := []string{strings.Split(m.selection, " ")[0]}
			return m, m.handleInstallPkgMsg(args)
		}

		m.selection = m.choices[m.cursor].title
		m.viewport.SetContent(m.renderContent())
		m.syncViewportScroll()
		return m, nil

	// Event Trigger Messages
	case messages.PkgOutput:
		m.SetPackages(msg.Output)
		m.viewport.SetContent(m.renderContent())
		return m, m.handleExecDoneMsg()
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}
