package panels

import (
	"github.com/ElitistNoob/pacdude/internal/tui/components/tabs"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Tab int

const (
	ViewAll Tab = iota
	Installed
	Updatable
)

type listKeyMap struct {
	Clear key.Binding
}

type ListPanel struct {
	List *list.Model
	Keys *listKeyMap
}

func (p *ListPanel) Init() tea.Cmd {
	return nil
}

func (p *ListPanel) Update(msg tea.Msg) (tabs.TabContent, tea.Cmd) {
	var cmd tea.Cmd
	*p.List, cmd = p.List.Update(msg)
	return p, cmd
}

func (p *ListPanel) View() string {
	return p.List.View()
}

func (p *ListPanel) SetSize(w, h int) {
	p.List.SetSize(w, h)
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		Clear: key.NewBinding(
			key.WithKeys("X"),
			key.WithHelp("X", "Clear"),
		),
	}
}

func NewListPanel(l list.Model) *ListPanel {
	keys := newListKeyMap()
	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			keys.Clear,
		}
	}
	l.SetShowTitle(false)
	// l.SetShowFilter(false)

	return &ListPanel{
		List: &l,
		Keys: keys,
	}
}

func (p *ListPanel) IsEmpty() bool {
	return len(p.List.Items()) == 0
}

func (p *ListPanel) SelectedItem() list.Item {
	return p.List.SelectedItem()
}

func (p *ListPanel) Query() string {
	return p.List.FilterValue()
}

func (p *ListPanel) ResetFilter() {
	p.List.ResetFilter()
}

func (p *ListPanel) ChangeFilterInput(s string) {
	p.List.FilterInput.Prompt = s
}

func (p *ListPanel) OpenSearchInput() {
	p.List.SetFilteringEnabled(true)
	p.List.SetShowFilter(true)
	p.List.FilterInput.Prompt = "Search Packages: "
	p.List.SetFilterState(list.Filtering)
	p.List.FilterInput.Focus()
}
