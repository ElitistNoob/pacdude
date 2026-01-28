package pkgs

import (
	"fmt"
	"os/exec"
	"strings"

	msg "github.com/ElitistNoob/pacdude/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

func execCmd(args []string) ([]byte, error) {
	cmd := exec.Command("pacman", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, fmt.Errorf("commmand failed : %w,\noutput: %s", err, output)
	}

	return output, err
}

func ExecWrapper(args []string) tea.Cmd {
	return func() tea.Msg {
		output, err := execCmd(args)
		return msg.PkgOutput{Output: output, Err: msg.ErrMsg{Err: err}}
	}
}

func parseOutput(output []byte) []pkg {
	lines := strings.Split(string(output), "\n")
	pkgs := make([]pkg, 0, len(lines)/2)

	for i := 0; i < len(lines)-1; i += 2 {
		title, desc := lines[i], lines[i+1]
		pkgs = append(pkgs, pkg{title: title, desc: desc})
	}

	return pkgs
}
