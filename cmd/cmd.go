package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Checks is binary of `rar` or `unrar` exists and executable
func Check(binary string) error {
	_, err := exec.LookPath(binary)
	if err != nil {
		return fmt.Errorf("binary %s not found or not executable", binary)
	}
	return nil
}

func Call(utility string, args []string) (string, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(utility, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("command failed: %v, stderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}
