package main

import (
	"fmt"
	"os"

	tui "github.com/ElitistNoob/pacdude/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	_, err := os.Stat("./debug.log")
	if err != nil {
		os.Create("debug.log")
	}
	os.Truncate("debug.log", 0)

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Printf("Couldn't open log files: %s\n", err)
	}
	defer f.Close()

	tui.Run()
}
