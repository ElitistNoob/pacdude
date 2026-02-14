package tui

import (
	"os/exec"

	"github.com/ElitistNoob/pacdude/internal/tui/choice"
	msg "github.com/ElitistNoob/pacdude/internal/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	stateReady state = iota
	stateConfirmInstall
	stateInstalling
	stateDone
	stateError
)

type outputMsg string
type errorMsg struct{ err error }

type model struct {
	state       state
	current     tea.Model
	previous    tea.Model
	selectedPkg string
	width       int
	height      int
	lastOutput  outputMsg
	lastErr     errorMsg
}

func newTuiModel() *model {
	return &model{
		current:    choice.InitialChoiceModel(),
		previous:   nil,
		lastOutput: "",
		lastErr:    errorMsg{},
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) installPkgCmd(pkg string) tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command(
			"sudo",
			"pacman",
			"-S",
			pkg,
			"--noconfirm",
			"--needed",
		)

		output, err := cmd.CombinedOutput()

		return msg.InstallResultMsg{
			Output: output,
			Err:    msg.ErrMsg{Err: err},
		}
	}
}
