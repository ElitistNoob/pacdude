package backend

import tea "github.com/charmbracelet/bubbletea"

func newResultMsg(output []byte, err error) tea.Cmd {
	return func() tea.Msg {
		return ResultMsg{
			Output: output,
			Err:    ErrMsg{Err: err},
		}
	}
}

func newInstallResultMsg(output []byte, err error) tea.Cmd {
	return func() tea.Msg {
		return InstallResultMsg{
			Result: ResultMsg{
				Output: output,
				Err:    ErrMsg{Err: err},
			},
		}
	}
}
