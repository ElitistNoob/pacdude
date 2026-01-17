package internal

import (
	"fmt"
	"os"

	"github.com/ElitistNoob/pacdude/internal/tui/root"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	p := tea.NewProgram(root.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
