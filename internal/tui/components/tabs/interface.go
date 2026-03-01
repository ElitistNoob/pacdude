package tabs

import tea "github.com/charmbracelet/bubbletea"

type TabContent interface {
	Init() tea.Cmd
	Update(tea.Msg) (TabContent, tea.Cmd)
	View() string
	SetSize(w, h int)
}
