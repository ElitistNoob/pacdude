package backend

import (
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type PacmanBackend struct{}

func (p PacmanBackend) ListInstalled() tea.Cmd {
	cmd := exec.Command("pacman", "-Qs")
	output, err := cmd.CombinedOutput()
	return newResultMsg(output, err)
}

func (p PacmanBackend) Search(query string) tea.Cmd {
	cmd := exec.Command("pacman", "-Ss", query)
	output, err := cmd.CombinedOutput()
	return newResultMsg(output, err)
}

func (p PacmanBackend) Install(pkg string) tea.Cmd {
	cmd := exec.Command("sudo", "pacman", "-S", pkg, "--noconfirm", "--needed")
	output, err := cmd.CombinedOutput()
	return newInstallResultMsg(output, err)
}

func (p PacmanBackend) Remove(pkg string) tea.Cmd {
	cmd := exec.Command("sudo", "pacman", "-Rns")
	output, err := cmd.CombinedOutput()
	return newResultMsg(output, err)
}

func (p PacmanBackend) ListUpgradable() tea.Cmd {
	cmd := exec.Command("pacman", "Qu")
	output, err := cmd.CombinedOutput()
	return newResultMsg(output, err)
}
