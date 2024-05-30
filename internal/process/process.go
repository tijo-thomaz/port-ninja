package process

import (
	"os/exec"
	"strconv"
)

func KillProcess(pid int) error {
	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	return cmd.Run()
}
