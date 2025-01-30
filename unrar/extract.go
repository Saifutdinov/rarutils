package unrar

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Saifutdinov/rarutils"
	commandline "github.com/Saifutdinov/rarutils/cmd"
	"github.com/Saifutdinov/rarutils/utils"
)

// Sets password to extract archive. "-pMyPassword".
func (a *Archive) SetPassword(password string) {
	a.password = password
}

// Sets overwrite mod to extract archive. "-o-" or "-o+".
func (a *Archive) SetOverwriteMode(overwrite bool) {
	a.notOverwrite = !overwrite
}

// Sets destination directory to extract, default "./".
func (a *Archive) SetDestination(destdir string) {
	a.destination = destdir
}

// Extracts and return files list
func (a *Archive) Extract() ([]Fileinfo, error) {
	return a.extract()
}

// Lists archive files list
func (a *Archive) List() ([]Fileinfo, error) {
	return a.list()
}

func (a *Archive) buildargs(action string) (args []string) {
	args = append(args, action)
	if a.password != "" {
		args = append(args, "-p"+a.password)
	}
	args = append(args, a.sourceFile)
	if action == actionExtract {

		if !a.notOverwrite {
			args = append(args, utils.Switch(a.notOverwrite, "-o-", "-o+"))
		}

		if a.destination == "" {
			a.setTempPath()
		}
		args = append(args, a.destination)
	}
	return
}

// Should I read file info from storage, or I can use unrar l?
func (a *Archive) extract() ([]Fileinfo, error) {
	args := a.buildargs(actionExtract)
	_, err := commandline.Call(rarutils.UnrarExeDefaultPath, args)
	if err != nil {
		return nil, err
	}
	return a.list()
}

func (a *Archive) list() ([]Fileinfo, error) {
	args := a.buildargs(actionList)
	output, err := commandline.Call(rarutils.UnrarExeDefaultPath, args)
	return parsefiles(output), err
}

func (a *Archive) setTempPath() string {
	a.destination = fmt.Sprintf("./%s_extracted_%d", a.fname(), time.Now().Nanosecond())
	return a.destination
}

func (a *Archive) fname() string {
	parts := strings.Split(a.sourceFile, "/")
	return parts[len(parts)-1]
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
