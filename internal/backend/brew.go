package backend

import (
	"encoding/json"
	"os/exec"
	"strings"

	msg "github.com/ElitistNoob/pacdude/internal/tui/messages"
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

func (p BrewBackend) ListInstalled() ResultMsg {
	cmd := exec.Command("brew", "info", "--json=v2", "--installed")
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionPackagesLoaded, err),
	}
}

func (p BrewBackend) ShowInfo(pkg string) (map[string]string, error) {
	result := make(map[string]string)
	cmd := exec.Command("brew", "info", pkg)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	// return ResultMsg{
	// 	Output:     output,
	// 	Err:        ErrMsg{Err: err},
	// 	ActionType: resolveAction(msg.ActionPackagesLoaded, err),
	// }
	result["Name"] = "Coming Soon"
	result["Reposity"] = "ElistNoob"
	result["Version"] = "0.0.0"
	result["Description"] = "Info not implemented for Brew"
	result["Packager"] = "ElitistNoob"
	return result, nil
}

func (p BrewBackend) ListAll() ResultMsg {
	cmd := exec.Command("brew", "info", "--json=v2", "--eval-all")
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionPackagesLoaded, err),
	}
}

func (p BrewBackend) Search(query string) ResultMsg {
	cmd := exec.Command("brew", "desc", "-s", query)
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionSearchLoaded, err),
	}
}

func (p BrewBackend) Install(pkg string) ResultMsg {
	cmd := exec.Command("brew", "install", pkg)
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionPackageInstalled, err),
	}
}

func (p BrewBackend) Remove(pkg string) ResultMsg {
	cmd := exec.Command("brew", "uninstall", pkg)
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionPackageRemoved, err),
	}
}

func (p BrewBackend) ListUpgradable() ResultMsg {
	_ = exec.Command("brew", "update").Run()

	cmd := exec.Command("brew", "outdated", "--verbose")
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionPackagesLoaded, err),
	}
}

func (p BrewBackend) UpdateAll() ResultMsg {
	cmd := exec.Command("brew", "upgrade")
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionUpdatedAll, err),
	}
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
			Name: strings.Join(strings.Split(split[0], ":"), ""),
			Desc: split[1],
		})
	}

	return pkgs
}
