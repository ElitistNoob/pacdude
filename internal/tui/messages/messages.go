package messages

import (
	tea "github.com/charmbracelet/bubbletea"
)

type OutputMsg []byte
type ErrMsg struct{ Err error }

type GoToPkgs struct{ Args []string }

type PkgOutput struct {
	Output OutputMsg
	Err    ErrMsg
}

func MsgHandler(args []string) tea.Cmd {
	return func() tea.Msg {
		return GoToPkgs{Args: args}
	}
}
