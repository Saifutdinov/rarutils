package main

import (
	"fmt"

	"github.com/Saifutdinov/rarutils/cmd"
	"github.com/Saifutdinov/rarutils/rar"
	"github.com/Saifutdinov/rarutils/unrar"
)

func init() {
	// Checking is rar utility exists and can be executed
	if err := cmd.Check(rar.RarExeFile); err != nil {
		fmt.Printf("Cannot run `rar`: %v \n", err)
	}
	// Checking is unrar utility exists and can be executed
	if err := cmd.Check(unrar.UnrarExeFile); err != nil {
		fmt.Printf("Cannot run `unrar`: %v \n", err)
	}
}

func main() {}
