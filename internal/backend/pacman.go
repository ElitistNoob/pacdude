package backend

import (
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type PacmanBackend struct{}

func (p PacmanBackend) ListInstalled() tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("pacman", "-Qs")
		output, err := cmd.CombinedOutput()
		return ListInstalledPackagesMsg{
			Output: output,
			Err:    ErrMsg{Err: err},
		}
	}
}

func (p PacmanBackend) Search(query string) tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("pacman", "-Ss", query)
		output, err := cmd.CombinedOutput()
		return SearchPacmanPackagesMsg{
			Output: output,
			Err:    ErrMsg{Err: err},
		}
	}
}

func (p PacmanBackend) Install(pkg string) tea.Cmd {
	cmd := exec.Command("sudo", "pacman", "-S", pkg)
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return InstallPackageResultMsg{Err: ErrMsg{Err: err}}
	})
}

func (p PacmanBackend) Remove(pkg string) tea.Cmd {
	cmd := exec.Command("sudo", "pacman", "-Rns", pkg)
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return RemovePackageResultMsg{Err: ErrMsg{Err: err}}
	})
}

func (p PacmanBackend) ListUpgradable() tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("checkupdates")
		output, err := cmd.CombinedOutput()
		return ListAvailableUpdatesMsg{
			Output: output,
			Err:    ErrMsg{Err: err},
		}
	}
}

func (p PacmanBackend) UpdateAll() tea.Cmd {
	cmd := exec.Command("sudo", "pacman", "-Syu")
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return UpdateAllMsg{Err: ErrMsg{Err: err}}
	})
}

func (p PacmanBackend) ParseOutput(output []byte) []Pkg {
	lines := strings.Split(string(output), "\n")
	pkgs := make([]Pkg, 0, len(lines)/2)

	for i := 0; i < len(lines)-1; i += 2 {
		title, desc := lines[i], lines[i+1]
		pkgs = append(pkgs, Pkg{Name: title, Desc: desc})
	}

	return pkgs
}
