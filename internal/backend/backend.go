package backend

import tea "github.com/charmbracelet/bubbletea"

type OutputMsg []byte
type ErrMsg struct{ Err error }

type ResultMsg struct {
	Output OutputMsg
	Err    ErrMsg
}

type InstallResultMsg struct {
	Result ResultMsg
}

type BackendInterface interface {
	ListInstalled() tea.Cmd
	Search(query string) tea.Cmd
	Install(pkg string) tea.Cmd
	Remove(pkg string) tea.Cmd
	ListUpgradable() tea.Cmd
}
