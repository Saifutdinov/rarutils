package unrar

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Saifutdinov/rarutils/utils"
)

type (
	Archive struct {
		Name     string
		Path     string
		Password string

		NotOverwrite bool
		Destination  string

		NotClear bool

		Files []*os.File
	}

	Fileinfo struct {
		Name string
		Size int64
	}
)

const (
	UnrarExeFile  = "/usr/local/bin/unrar"
	ActionExtract = "x"
	ActionList    = "lt"
)

func (a Archive) Extract() {
	a.extract()
}

func (a Archive) List() []Fileinfo {
	return a.list()
}

func (a Archive) Stream(keepAfterReturn bool) {
	a.extract()

	if !keepAfterReturn {
		os.RemoveAll(a.Destination)
	}
}

func (a Archive) buildargs(action string) (args []string) {
	args = append(args, action)
	if a.Password != "" {
		args = append(args, "-p"+a.Password)
	}

	args = append(args, a.Path)

	if action == ActionExtract {

		if !a.NotOverwrite {
			args = append(args, utils.Switch(a.NotOverwrite, "-o-", "-o+"))
		}

		if a.Destination == "" {
			a.Destination = a.temppath()
		}
		args = append(args, a.Destination)
	}
	return
}

func (a Archive) extract() {
	args := a.buildargs(ActionExtract)
	utils.CMD(UnrarExeFile, args)
}

func (a Archive) list() []Fileinfo {
	args := a.buildargs(ActionList)
	output, _ := utils.CMD(UnrarExeFile, args)

	return readfiles(output)
}

func (a Archive) temppath() string {
	return fmt.Sprintf("./%s_extracted_%d", a.Path, time.Now().Nanosecond())
}

func readfiles(output string) (files []Fileinfo) {
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		fmt.Println(line)
		// if strings.Contains(line, ".....") {
		// 	parts := strings.Fields(line)
		// 	if len(parts) >= 4 {
		// 		files = append(files, Fileinfo{
		// 			Name: parts[len(parts)-1],
		// 			Size: parseSize(parts[0]),
		// 		})
		// 	}
		// }
	}

	return files
}

// func parseSize(sizeStr string) int64 {
// 	size, err := strconv.ParseInt(sizeStr, 2, 64)
// 	if err != nil {
// 		return 0
// 	}
// 	return size
// }
