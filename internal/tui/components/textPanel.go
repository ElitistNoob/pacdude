package panels

import (
	"github.com/ElitistNoob/pacdude/internal/tui/components/tabs"
	tea "github.com/charmbracelet/bubbletea"
)

type TextPanel struct {
	Text string
}

func (p *TextPanel) Init() tea.Cmd {
	return nil
}

func (p *TextPanel) Update(msg tea.Msg) (tabs.TabContent, tea.Cmd) {
	return p, nil
}

func (p *TextPanel) View() string {
	return p.Text
}

func (p *TextPanel) SetSize(w, h int) {}

func (p *TextPanel) NewTextPanel(text string) *TextPanel {
	return &TextPanel{
		Text: text,
	}
}

//
// func (p *TextPanel) SelectedItem() list.Item {
// 	return p.Text.SelectedItem()
// }
