package backend

import (
	"os/exec"
	"strings"

	msg "github.com/ElitistNoob/pacdude/internal/tui/messages"
)

type PacmanBackend struct{}

func (p PacmanBackend) ListInstalled() ResultMsg {
	cmd := exec.Command("pacman", "-Qs")
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionInstalledLoaded, err),
	}
}

func (p PacmanBackend) Search(query string) ResultMsg {
	cmd := exec.Command("pacman", "-Ss", query)
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionSearchLoaded, err),
	}
}

func (p PacmanBackend) Install(pkg string) ResultMsg {
	cmd := exec.Command("sudo", "pacman", "-S", pkg, "--noconfirm", "--needed")
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionPackageInstalled, err),
	}
}

func (p PacmanBackend) Remove(pkg string) ResultMsg {
	cmd := exec.Command("sudo", "pacman", "-Rns", pkg, "--noconfirm")
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionPackageRemoved, err),
	}
}

func (p PacmanBackend) ListUpgradable() ResultMsg {
	cmd := exec.Command("checkupdates")
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionUpdatesLoaded, err),
	}
}

func (p PacmanBackend) UpdateAll() ResultMsg {
	cmd := exec.Command("sudo", "pacman", "-Syu")
	output, err := cmd.CombinedOutput()
	return ResultMsg{

		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionUpdatedAll, err),
	}
}

func (p PacmanBackend) ParseOutput(output []byte) []Pkg {
	lines := strings.Split(string(output), "\n")
	pkgs := make([]Pkg, 0, len(lines)/2)

	for i := 0; i < len(lines)-1; i += 2 {
		title, desc := lines[i], lines[i+1]
		pkgs = append(pkgs, Pkg{Name: title, Desc: strings.TrimSpace(desc)})
	}

	return pkgs
}
