package rar

import (
	"fmt"
	"os"
	"strings"

	"github.com/Saifutdinov/utils"
)

type (
	Archive struct {
		// how to name file. No ".rar" needed in the end.
		Name string
		// where save file
		DestinationDir string

		// save as solid
		Solid bool

		//Directory of files. Example - /path/to/directory
		SourceDir string
		//File pattern of files. Example - /path/to/files/*.pdf
		FilePattern string
		// List of file paths. Example - [/path/to/file1.pdf, /path/to/file2.pdf, /path/to/file3.pdf, ...]
		Files []string

		// Compression. Example - m0 - m5. Default empty.
		Compression string

		// Volumes. Example - v10MB. Default empty.
		Volumes string

		// Password
		Password string
	}
)

const (
	NoCompression   = ""
	CompressionLVL0 = "m0"
	CompressionLVL1 = "m1"
	CompressionLVL2 = "m2"
	CompressionLVL3 = "m3"
	CompressionLVL4 = "m4"
	CompressionLVL5 = "m5"
)

const (
	RarExeFile        = "/usr/local/bin/rar"
	FilesListFileName = "rarfileslist*"
)

// Store your rar file to path
func (a Archive) Save() {
	a.savefile()
}

// Creates file with params, returns you []byte to force download or send by email
// and then removes file from path (you need just []byte)
func (a Archive) Stream(keepAfterReturn bool) []byte {
	a.savefile()
	file, _ := os.ReadFile(a.filename())

	if !keepAfterReturn {
		os.Remove(a.filename())
	}
	return file
}

func (a Archive) filename() string {
	if a.DestinationDir == "" {
		a.DestinationDir = "."
	}
	return fmt.Sprintf("%s/%s.rar", a.DestinationDir, a.Name)
}

func (a Archive) buildargs() (args []string, tempfile string) {
	args = append(args, "a", a.filename())
	if a.Solid {
		args = append(args, "-s")
	}
	if a.Compression != "" {
		args = append(args, a.Compression)
	}
	if a.Volumes != "" {
		args = append(args, "-v"+a.Volumes)
	}
	source, tempfile := a.source()
	args = append(args, source...)

	if a.Password != "" {
		args = append(args, "-p"+a.Password)
	}
	return
}

func (a Archive) source() (source []string, tempfile string) {
	if a.SourceDir != "" {
		source = append(source, "-r", a.SourceDir)
		return
	}
	if a.FilePattern != "" {
		source = append(source, a.FilePattern)
	}

	tempfile = createFilesList(a.Files)
	source = append(source, "@"+tempfile)
	return
}

// executes command to create rar archive file
func (a Archive) savefile() {
	args, tempfile := a.buildargs()
	utils.CMD(RarExeFile, args)

	if tempfile != "" {
		os.RemoveAll(tempfile)
	}
}

// creates tmp file to save as argument "@fileslist.txt" to create archive file
func createFilesList(fs []string) string {
	tempFile, _ := os.CreateTemp("", FilesListFileName)
	defer tempFile.Close()
	fileslist := strings.Join(fs, "\n")
	tempFile.WriteString(fileslist)
	return tempFile.Name()
}
