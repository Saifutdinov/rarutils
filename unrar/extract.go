package unrar

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Saifutdinov/rarutils"
	commandline "github.com/Saifutdinov/rarutils/cmd"
	"github.com/Saifutdinov/rarutils/utils"
)

func (a *Archive) SetPassword(password string) {
	a.Password = password
}

func (a *Archive) SetOverwriteMode(overwrite bool) {
	a.NotOverwrite = !overwrite
}

func (a *Archive) SetDestination(destdir string) {
	a.Destination = destdir
}

func (a *Archive) Extract() ([]Fileinfo, error) {
	return a.extract()
}

func (a *Archive) List() ([]Fileinfo, error) {
	return a.list()
}

func (a *Archive) Stream(keepAfterReturn bool) {
	a.extract()

	if !keepAfterReturn {
		os.RemoveAll(a.Destination)
	}
}

func (a *Archive) buildargs(action string) (args []string) {
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
			a.setTempPath()
		}
		args = append(args, a.Destination)
	}
	return
}

// Should I read file info from storage, or I can use unrar l?
func (a *Archive) extract() ([]Fileinfo, error) {
	args := a.buildargs(ActionExtract)
	_, err := commandline.Call(rarutils.UnrarExeDefaultPath, args)
	if err != nil {
		return nil, err
	}
	return a.list()
}

func (a *Archive) list() ([]Fileinfo, error) {
	args := a.buildargs(ActionList)
	output, err := commandline.Call(rarutils.UnrarExeDefaultPath, args)
	return parsefiles(output), err
}

func (a *Archive) setTempPath() string {
	a.Destination = fmt.Sprintf("./%s_extracted_%d", a.SourceFile, time.Now().Nanosecond())
	return a.Destination
}

// func readfiles(destdir string) ([]Fileinfo, error) {
// 	files, err := os.ReadDir(destdir)
// 	if err != nil {
// 		return nil, err
// 	}
// }

// TODO: upgrade to unrar lt info
func parsefiles(output string) (files []Fileinfo) {
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
