package backend

type OutputMsg []byte
type ErrMsg struct{ Err error }

type ResultMsg struct {
	Output OutputMsg
	Err    ErrMsg
}

type BackendInterface interface {
	ListInstalled() string
	Search(query string) ResultMsg
	Install(pkg string) ResultMsg
	Remove(pkg string) ResultMsg
	ListUpgradable() ResultMsg
}
