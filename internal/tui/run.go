package tui

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ElitistNoob/pacdude/internal/app"
	backendselector "github.com/ElitistNoob/pacdude/internal/tui/screens/backendSelector"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	cmd := exec.Command("sudo", "-v")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("sudo authentication failed", err)
		os.Exit(1)
	}

	p := tea.NewProgram(app.AppModel{
		Current: backendselector.NewModel(),
	},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
