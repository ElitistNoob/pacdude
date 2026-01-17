package app

import (
	"fmt"
	"os/exec"
)

func InstalledPackages() (string, error) {
	cmd := exec.Command("pacman", "-Qs")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command failed: %w\noutput: %s", err, output)
	}

	return fmt.Sprint(output), nil
}
