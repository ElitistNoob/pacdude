package backend

import (
	"encoding/json"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type BrewPackage struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Full string `json:"full_name"`
}

type BrewRoot struct {
	Formulae []BrewPackage `json:"formulae"`
}

type BrewBackend struct{}

func (p BrewBackend) ListInstalled() tea.Cmd {
	cmd := exec.Command("brew", "info", "--json=v2", "--installed")
	output, err := cmd.CombinedOutput()
	return func() tea.Msg {
		return ListInstalledPackagesMsg{
			Output: output,
			Err:    ErrMsg{Err: err},
		}
	}
}

func (p BrewBackend) Search(query string) tea.Cmd {
	cmd := exec.Command("brew", "desc", "-s", query)
	output, err := cmd.CombinedOutput()
	return func() tea.Msg {
		return SearchPacmanPackagesMsg{
			Output: output,
			Err:    ErrMsg{Err: err},
		}
	}
}

func (p BrewBackend) Install(pkg string) tea.Cmd {
	cmd := exec.Command("brew", "install", pkg)
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return InstallPackageResultMsg{Err: ErrMsg{Err: err}}
	})
}

func (p BrewBackend) Remove(pkg string) tea.Cmd {
	cmd := exec.Command("brew", "uninstall", pkg)
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return RemovePackageResultMsg{Err: ErrMsg{Err: err}}
	})
}

func (p BrewBackend) ListUpgradable() tea.Cmd {
	_ = exec.Command("brew", "update").Run()

	cmd := exec.Command("brew", "outdated", "--verbose")
	output, err := cmd.CombinedOutput()
	return func() tea.Msg {
		return ListAvailableUpdatesMsg{
			Output: output,
			Err:    ErrMsg{Err: err},
		}
	}
}

func (p BrewBackend) UpdateAll() tea.Cmd {
	cmd := exec.Command("brew", "upgrade")
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return UpdateAllMsg{Err: ErrMsg{Err: err}}
	})
}

func (p BrewBackend) ParseOutput(output []byte) []Pkg {
	var data BrewRoot
	if err := json.Unmarshal(output, &data); err == nil && len(data.Formulae) > 0 {
		pkgs := make([]Pkg, 0, len(data.Formulae))
		for _, f := range data.Formulae {
			pkgs = append(pkgs, Pkg{
				Name: f.Name,
				Desc: f.Desc,
			})
		}
		return pkgs
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	pkgs := make([]Pkg, 0, len(lines))
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if line == "" || line == "==> Formulae" || line == "==> Casks" {
			continue
		}
		split := strings.SplitN(line, " ", 2)
		pkgs = append(pkgs, Pkg{
			Name: split[0],
			Desc: strings.Join(strings.Split(split[1], ":"), ""),
		})
	}

	return pkgs
}
