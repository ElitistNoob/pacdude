package backend

import "os/exec"

func newResultMsg(cmd exec.Cmd) ResultMsg {
	output, err := cmd.CombinedOutput()
	return ResultMsg{
		Output: output,
		Err:    ErrMsg{Err: err},
	}
}
