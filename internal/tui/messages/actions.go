package messages

type ActionType int

const (
	ActionNone ActionType = iota
	ActionInstalledLoaded
	ActionUpdatesLoaded
	ActionSearchLoaded
	ActionPackageInstalled
	ActionPackageRemoved
	ActionUpdatedAll
	ActionError
)

type ActionMsg struct {
	Type    ActionType
	Payload any
	Err     error
}

