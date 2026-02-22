package tabs

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

type Tab int

const (
	Installed Tab = iota
	Updatable
)

type listKeyMap struct {
	InstalledPackage key.Binding
	Install          key.Binding
	Updatable        key.Binding
	UpdateAll        key.Binding
	Uninstall        key.Binding
	NextTab          key.Binding
	PrevTab          key.Binding
}

type TabsModel struct {
	Tabs  []list.Model
	Index Tab
	Keys  *listKeyMap
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		InstalledPackage: key.NewBinding(
			key.WithKeys("I"),
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
	}
}

func NewTabsModel() *TabsModel {
	tabsTitles := []string{
		"Installed (I)",
		"Available Updates (U)",
	}
	tabs := make([]list.Model, len(tabsTitles))
	listKey := newListKeyMap()
	for i := range tabs {
		l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
		l.Title = tabsTitles[i]
		l.SetShowTitle(true)
		l.SetShowStatusBar(true)
		l.AdditionalFullHelpKeys = func() []key.Binding {
			return []key.Binding{
				listKey.InstalledPackage,
				listKey.Install,
				listKey.Updatable,
				listKey.UpdateAll,
				listKey.Uninstall,
				listKey.NextTab,
				listKey.PrevTab,
			}
		}

		tabs[i] = l
	}

	return &TabsModel{
		Index: 0,
		Tabs:  tabs,
		Keys:  listKey,
	}
}

func (m *TabsModel) SetSize(w, h int) {
	for i := range m.Tabs {
		m.Tabs[i].SetSize(w, h)
	}
}

func (m *TabsModel) Active() *list.Model {
	return &m.Tabs[m.Index]
}

func (m *TabsModel) SetActive(l list.Model) {
	m.Tabs[m.Index] = l
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


func (m *TabsModel) NextTab() {
	l := len(m.Tabs)
	m.Index = Tab((int(m.Index) + 1) % l)
}

func (m *TabsModel) PrevTab() {
	l := len(m.Tabs)
	m.Index = Tab((int(m.Index) - 1 + l) % l)
}
