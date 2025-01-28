package rarutils

import (
	"fmt"

	"github.com/Saifutdinov/rarutils/cmd"
)

var (
	// Default path to rar/unrar binaries
	RarExeDefaultPath   = "/usr/local/bin/rar"
	UnrarExeDefaultPath = "/usr/local/bin/unrar"
)

func init() {
	// Checking is rar utility exists and can be executed
	if err := cmd.Check(RarExeDefaultPath); err != nil {
		fmt.Printf("Cannot run `rar`: %v \n", err)
	}
	// Checking is unrar utility exists and can be executed
	if err := cmd.Check(UnrarExeDefaultPath); err != nil {
		fmt.Printf("Cannot run `unrar`: %v \n", err)
	}

	// archive := unrar.NewArchive("path")

	// archive.SetDestination("/path/to/extract")
	// // -o+ / -o-
	// archive.SetOverwriteMode(true)
	// // -pMySecretPassword
	// archive.SetPassword("MySecretPassword")

}

// Sets `rar` exe file if not located in default path
func SetRarPath(path string) {
	RarExeDefaultPath = path
}

// Sets `unrar` exe file if not located in default path
func SetUnarPath(path string) {
	UnrarExeDefaultPath = path
}
