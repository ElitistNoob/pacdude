package backend

import (
	"os/exec"
	"strings"

	msg "github.com/ElitistNoob/pacdude/internal/tui/messages"
	"github.com/ElitistNoob/pacdude/internal/tui/styles"
)

type PacmanBackend struct{}

func (p PacmanBackend) ListInstalled() ResultMsg {
	cmd := exec.Command("pacman", "-Qs")
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionPackagesLoaded, err),
	}
}

func (p PacmanBackend) ShowInfo(pkg string) (map[string]string, error) {
	result := make(map[string]string)

	cmd := exec.Command("pacman", "-Si", pkg)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var currentkey string
	lines := strings.SplitSeq(string(output), "\n")
	for line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
			if currentkey != "" {
				result[currentkey] += " " + strings.TrimSpace(line)
			}
			continue

		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		result[key] = value
		currentkey = key

	}

	return result, nil
}

func (p PacmanBackend) ListAll() ResultMsg {
	cmd := exec.Command("pacman", "-Ss")
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output:     output,
		Err:        ErrMsg{Err: err},
		ActionType: resolveAction(msg.ActionPackagesLoaded, err),
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
		ActionType: resolveAction(msg.ActionPackagesLoaded, err),
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
		parts := strings.Split(title, " ")
		if len(parts) == 3 {
			parts[2] = styles.InstalledPackage.Render(parts[2])
		}

		title = strings.Join(parts, " ")
		pkgs = append(pkgs, Pkg{Name: title, Desc: strings.TrimSpace(desc)})
	}

	return pkgs
}
