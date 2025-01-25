package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func CheckBinary(binary string) error {
	_, err := exec.LookPath(binary)
	if err != nil {
		return fmt.Errorf("binary %s not found or not executable", binary)
	}
	return nil
}

func CMD(utility string, args []string) (string, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(utility, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("command failed: %v, stderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}
