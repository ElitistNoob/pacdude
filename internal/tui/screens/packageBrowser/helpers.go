package packagebrowser

import (
	"strings"
)

func (m *PackageBrowserModel) parseOutput(output []byte) []pkg {
	lines := strings.Split(string(output), "\n")
	pkgs := make([]pkg, 0, len(lines)/2)

	for i := 0; i < len(lines)-1; i += 2 {
		title, desc := lines[i], lines[i+1]
		pkgs = append(pkgs, pkg{title: title, desc: desc})
	}

	return pkgs
}
