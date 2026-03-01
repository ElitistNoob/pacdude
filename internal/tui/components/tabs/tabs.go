package tabs

type TabsModel struct {
	Tabs       []string
	TabContent []TabContent
	Index      int
}

func NewTabsModel(tabs []string, content []TabContent) *TabsModel {
	return &TabsModel{
		Tabs:       tabs,
		TabContent: content,
		Index:      0,
	}
}

func (m *TabsModel) SetSize(w, h int) {
	for i := range m.Tabs {
		m.TabContent[i].SetSize(w, h)
	}
}

func (m *TabsModel) Active() TabContent {
	return m.TabContent[m.Index]
}

func (m *TabsModel) NextTab() {
	l := len(m.Tabs)
	m.Index = (m.Index + 1) % l
}

func (m *TabsModel) PrevTab() {
	l := len(m.Tabs)
	m.Index = (m.Index - 1 + l) % l
}
