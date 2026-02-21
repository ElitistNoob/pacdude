package backend

import (
	"github.com/ElitistNoob/pacdude/internal/tui/messages"
)

type Pkg struct {
	Name, Desc string
}

func (p Pkg) Title() string       { return p.Name }
func (p Pkg) Description() string { return p.Desc }
func (p Pkg) FilterValue() string { return p.Name }

type OutputMsg []byte
type ErrMsg struct{ Err error }

type ResultMsg struct {
	Output     OutputMsg
	Err        ErrMsg
	ActionType messages.ActionType
}

type BackendInterface interface {
	ListInstalled() ResultMsg
	Search(query string) ResultMsg
	Install(pkg string) ResultMsg
	Remove(pkg string) ResultMsg
	ListUpgradable() ResultMsg
	UpdateAll() ResultMsg
	ParseOutput(output []byte) []Pkg
}
