package backend

import "github.com/ElitistNoob/pacdude/internal/tui/messages"

func resolveAction(action messages.ActionType, err error) messages.ActionType {
	actionType := action
	if err != nil {
		actionType = messages.ActionError
	}

	return actionType
}
