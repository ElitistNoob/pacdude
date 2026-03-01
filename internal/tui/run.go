package tui

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ElitistNoob/pacdude/internal/app"
	"github.com/ElitistNoob/pacdude/internal/backend"
	packagebrowser "github.com/ElitistNoob/pacdude/internal/tui/screens/packageBrowser"
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
		Current: packagebrowser.NewModel(backend.PacmanBackend{}, 0),
	},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
