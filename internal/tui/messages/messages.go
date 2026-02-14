package messages

// Core Msg
type OutputMsg []byte
type ErrMsg struct{ Err error }

// Package screen related msgs
type GoToPkgsMsg struct{ Args []string }
type PkgOutput struct {
	Output OutputMsg
	Err    ErrMsg
}

// Install Process related msgs
type InstallPkgMsg struct{ Args []string }
type InstallResultMsg struct {
	Output OutputMsg
	Err    ErrMsg
}

// State messages
type ExecDoneMsg struct{}
