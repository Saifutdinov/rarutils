package unrar

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	commandline "github.com/Saifutdinov/rarutils/cmd"
	"github.com/Saifutdinov/rarutils/utils"
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
	args = append(args, a.SourceFile)
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
	commandline.Call(UnrarExeFile, args)
}

func (a Archive) list() []Fileinfo {
	args := a.buildargs(ActionList)
	output, _ := commandline.Call(UnrarExeFile, args)

	return readfiles(output)
}

func (a Archive) temppath() string {
	return fmt.Sprintf("./%s_extracted_%d", a.SourceFile, time.Now().Nanosecond())
}

// TODO: upgrade to unrar lt info
func readfiles(output string) (files []Fileinfo) {
	lines := strings.Split(output, "\n")
	filesline := 0
	for i, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		if parts[0] == "Attributes" {
			filesline = i + 1
		}
		if filesline > 0 && i > filesline {
			files = append(files, Fileinfo{
				Name: parts[len(parts)-1],
				Size: parseSize(parts[1]),
			})
		}
		if filesline > 0 && parts[0] == "-----------" && i > filesline {
			break
		}
	}

	return files
}

func parseSize(sizeStr string) int64 {
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		return 0
	}
	return size
}
