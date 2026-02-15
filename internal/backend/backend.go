package backend

import tea "github.com/charmbracelet/bubbletea"

type Pkg struct {
	Name, Desc string
}

func (p Pkg) Title() string       { return p.Name }
func (p Pkg) Description() string { return p.Desc }
func (p Pkg) FilterValue() string { return p.Name }

type OutputMsg []byte
type ErrMsg struct{ Err error }

type ListInstalledPackagesMsg struct {
	Output OutputMsg
	Err    ErrMsg
}

type InstallPackageResultMsg struct {
	Err ErrMsg
}

type RemovePackageResultMsg struct {
	Err ErrMsg
}

type SearchPacmanPackagesMsg struct {
	Output OutputMsg
	Err    ErrMsg
}

type ListAvailableUpdatesMsg struct {
	Output OutputMsg
	Err    ErrMsg
}

type UpdateAllMsg struct {
	Err ErrMsg
}

type BackendInterface interface {
	ListInstalled() tea.Cmd
	Search(query string) tea.Cmd
	Install(pkg string) tea.Cmd
	Remove(pkg string) tea.Cmd
	ListUpgradable() tea.Cmd
	UpdateAll() tea.Cmd
	ParseOutput(output []byte) []Pkg
}
