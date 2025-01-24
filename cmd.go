package rarutils

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckBinary(binary string) error {
	_, err := exec.LookPath(binary)
	if err != nil {
		return fmt.Errorf("binary %s not found or not executable", binary)
	}
	return nil
}

func CMD(utility string, args []string) {
	cmd := exec.Command(utility, args...)

	fmt.Println(cmd.String())

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
