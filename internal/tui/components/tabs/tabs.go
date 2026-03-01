package tabs

import (
	"github.com/ElitistNoob/pacdude/internal/backend"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

type Tab int

const (
	ViewAll Tab = iota
	Installed
	Updatable
)

type listKeyMap struct {
	ViewAll          key.Binding
	InstalledPackage key.Binding
	Install          key.Binding
	Updatable        key.Binding
	UpdateAll        key.Binding
	Uninstall        key.Binding
	NextTab          key.Binding
	PrevTab          key.Binding
	Clear            key.Binding
}

type TabsModel struct {
	Tabs  []string
	Lists []list.Model
	Index Tab
	Keys  *listKeyMap
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		ViewAll: key.NewBinding(
			key.WithKeys("A"),
			key.WithHelp("A", "All Packages"),
		),
		InstalledPackage: key.NewBinding(
			key.WithKeys("I"),
			key.WithHelp("I", "Installed Packages"),
		),
		Install: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Install"),
		),
		Updatable: key.NewBinding(
			key.WithKeys("U"),
			key.WithHelp("U", "Available Updates"),
		),
		Uninstall: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "Uninstall"),
		),
		NextTab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("Tab", "Next Tab"),
		),
		PrevTab: key.NewBinding(
			key.WithKeys("shift+tab"),
			key.WithHelp("Shift+Tab", "Previous Tab"),
		),
		Clear: key.NewBinding(
			key.WithKeys("X"),
			key.WithHelp("X", "Clear"),
		),
	}
}

func NewTabsModel() *TabsModel {
	tabsTitles := []string{
		"All (A)",
		"Installed (I)",
		"Updates (U)",
	}
	lists := make([]list.Model, len(tabsTitles))
	listKey := newListKeyMap()
	for i := range tabsTitles {
		l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
		l.SetShowTitle(false)
		// l.SetShowHelp(false)
		l.AdditionalFullHelpKeys = func() []key.Binding {
			return []key.Binding{
				listKey.ViewAll,
				listKey.InstalledPackage,
				listKey.Install,
				listKey.Updatable,
				listKey.UpdateAll,
				listKey.Uninstall,
				listKey.NextTab,
				listKey.PrevTab,
				listKey.Clear,
			}
		}

		lists[i] = l
	}

	return &TabsModel{
		Index: 0,
		Tabs:  tabsTitles,
		Lists: lists,
		Keys:  listKey,
	}
}

func (m *TabsModel) SetSize(w, h int) {
	for i := range m.Tabs {
		m.Lists[i].SetSize(w, h)
	}
}

func (m *TabsModel) Active() *list.Model {
	return &m.Lists[m.Index]
}

func (m *TabsModel) SetActive(l list.Model) {
	m.Lists[m.Index] = l
}

func (m *TabsModel) IsActiveEmpty() bool {
	return len(m.Active().Items()) == 0
}

func (m *TabsModel) SelectedItem() list.Item {
	return m.Active().SelectedItem()
}

func (m *TabsModel) Query() string {
	return m.Active().FilterValue()
}

func (m *TabsModel) ResetFilter() {
	m.Active().ResetFilter()
}

func (m *TabsModel) ChangeFilterInput(s string) {
	m.Active().FilterInput.Prompt = s
}

func (m *TabsModel) NextTab() {
	l := len(m.Tabs)
	m.Index = Tab((int(m.Index) + 1) % l)
}

func (m *TabsModel) PrevTab() {
	l := len(m.Tabs)
	m.Index = Tab((int(m.Index) - 1 + l) % l)
}

func (m *TabsModel) OpenSearchInput() {
	m.Active().SetFilteringEnabled(true)
	m.Active().SetShowFilter(true)
	m.Active().FilterInput.Prompt = "Search Packages: "
	m.Active().SetFilterState(list.Filtering)
	m.Active().FilterInput.Focus()
}

func (m *TabsModel) CallAction(b backend.BackendInterface) backend.ResultMsg {
	var result backend.ResultMsg
	switch m.Index {
	case Installed:
		if m.IsActiveEmpty() {
			result = b.ListInstalled()
		}
	case Updatable:
		if m.IsActiveEmpty() {
			result = b.ListUpgradable()
		}
	case ViewAll:
		if m.IsActiveEmpty() {
			result = b.ListAll()
		}
	}

	return result
}
