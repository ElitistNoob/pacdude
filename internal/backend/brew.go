package backend

import (
	"encoding/json"
	"os/exec"

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
	cmd := exec.Command("brew", "info", "--json=v2", "-installed")
	output, err := cmd.CombinedOutput()
	return func() tea.Msg {
		return ListInstalledPackagesMsg{
			Output: output,
			Err:    ErrMsg{Err: err},
		}
	}
}

func (p BrewBackend) Search(query string) tea.Cmd {
	cmd := exec.Command("brew", "search", query)
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

	cmd := exec.Command("brew", "outdated")
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
	json.Unmarshal(output, &data)

	pkgs := make([]Pkg, 0, len(data.Formulae))
	for i, f := range pkgs {
		pkgs[i] = Pkg{
			Name: f.Name,
			Desc: f.Desc,
		}
	}

	return pkgs
}
