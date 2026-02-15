package packagebrowser

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func parseOutput(output []byte) []pkg {
	lines := strings.Split(string(output), "\n")
	pkgs := make([]pkg, 0, len(lines)/2)

	for i := 0; i < len(lines)-1; i += 2 {
		title, desc := lines[i], lines[i+1]
		pkgs = append(pkgs, pkg{title: title, desc: desc})
	}

	return pkgs
}

func (m *PackageBrowserModel) setListItems(output []byte) tea.Cmd {
	return func() tea.Msg {
		o := parseOutput(output)
		items := make([]list.Item, len(o))
		for i, v := range o {
			items[i] = v
		}
		return m.list.SetItems(items)
	}
}
