package rarutils

import (
	"fmt"

	"github.com/Saifutdinov/rarutils/cmd"
)

var (
	//
	showLogs = true

	// Default path to rar/unrar binaries
	RarExeDefaultPath   = "/usr/local/bin/rar"
	UnrarExeDefaultPath = "/usr/local/bin/unrar"
)

func init() {
	if showLogs {
		// Checking is rar utility exists and can be executed
		if err := cmd.Check(RarExeDefaultPath); err != nil {
			fmt.Printf("Be careful! %v \n", err)
		}
		// Checking is unrar utility exists and can be executed
		if err := cmd.Check(UnrarExeDefaultPath); err != nil {
			fmt.Printf("Be careful! %v \n", err)
		}
	}
}

func CheckBinaries(check bool) {
	showLogs = check
}

// Sets `rar` exe file if not located in default path
func SetRarPath(path string) {
	RarExeDefaultPath = path
}

// Sets `unrar` exe file if not located in default path
func SetUnrarPath(path string) {
	UnrarExeDefaultPath = path
}
