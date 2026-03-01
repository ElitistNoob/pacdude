package backend

import (
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
)

type Pkg struct {
	Name, Desc string
}

type PkgInfo struct {
	Repository,
	Name,
	Version,
	Description,
	Packager string
}

func (p Pkg) Title() string       { return p.Name }
func (p Pkg) Description() string { return p.Desc }
func (p Pkg) FilterValue() string { return p.Name }

type (
	OutputMsg []byte
	ErrMsg    struct{ Err error }
)

type ResultMsg struct {
	Output     OutputMsg
	Err        ErrMsg
	ActionType messages.ActionType
}

type BackendInterface interface {
	ListInstalled() ResultMsg
	ShowInfo(pkg string) (map[string]string, error)
	Search(query string) ResultMsg
	Install(pkg string) ResultMsg
	Remove(pkg string) ResultMsg
	ListUpgradable() ResultMsg
	UpdateAll() ResultMsg
	ListAll() ResultMsg
	ParseOutput(output []byte) []Pkg
}
