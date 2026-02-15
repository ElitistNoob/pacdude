package backend

import (
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type FlatpakBackend struct{}

func (p FlatpakBackend) ListInstalled() tea.Cmd {
	cmd := exec.Command("flatpak", "list", "--columns=name,application,description")
	output, err := cmd.CombinedOutput()
	return func() tea.Msg {
		return ListInstalledPackagesMsg{
			Output: output,
			Err:    ErrMsg{Err: err},
		}
	}
}

func (p FlatpakBackend) Search(query string) tea.Cmd {
	cmd := exec.Command("flatpak", "search", "--columns=name,application,description", query)
	output, err := cmd.CombinedOutput()
	return func() tea.Msg {
		return SearchPacmanPackagesMsg{
			Output: output,
			Err:    ErrMsg{Err: err},
		}
	}
}

func (p FlatpakBackend) Install(pkg string) tea.Cmd {
	// NEED TO CHANGE LATER THIS WON'T WORK -> Pkg needs to be the ID not the name
	cmd := exec.Command("flatpak", "install", "--user", "-y", pkg)
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return InstallPackageResultMsg{Err: ErrMsg{Err: err}}
	})
}

func (p FlatpakBackend) Remove(pkg string) tea.Cmd {
	// NEED TO CHANGE LATER THIS WON'T WORK -> Pkg needs to be the ID not the name
	cmd := exec.Command("flatpak", "uninstall", "--user", "-y", pkg)
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return RemovePackageResultMsg{Err: ErrMsg{Err: err}}
	})
}

func (p FlatpakBackend) ListUpgradable() tea.Cmd {
	_ = exec.Command("flatpak", "update", "--appstream").Run()

	cmd := exec.Command("flatpak", "remote-ls", "--updates")
	output, err := cmd.CombinedOutput()
	return func() tea.Msg {
		return ListAvailableUpdatesMsg{
			Output: output,
			Err:    ErrMsg{Err: err},
		}
	}
}

func (p FlatpakBackend) UpdateAll() tea.Cmd {
	cmd := exec.Command("flatpak", "update", "-y")
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return UpdateAllMsg{Err: ErrMsg{Err: err}}
	})
}

func (p FlatpakBackend) ParseOutput(output []byte) []Pkg {
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	pkgs := make([]Pkg, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, "\t")
		if len(parts) >= 3 {
			pkgs = append(pkgs, Pkg{
				Name: fmt.Sprintf("%s (%s)", parts[0], parts[1]),
				Desc: parts[2],
			})
		}
	}

	return pkgs
}
