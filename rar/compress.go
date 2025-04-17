package rar

import (
	"fmt"
	"os"
	"strings"

	"github.com/Saifutdinov/rarutils"
	"github.com/Saifutdinov/rarutils/cmd"
)

const filesListFileName = "rarfileslist*.lst"

// Sets source directory to compress
func (a *ArchiveFile) SetSourceDir(dir string) {
	a.sourceDir = dir
}

// Sets destination directory to store archive
func (a *ArchiveFile) SetDestinationDir(dir string) {
	a.destinationDir = dir
}

// Sets pattern of file to compress
func (a *ArchiveFile) SetFilePattern(pattern string) {
	a.filePattern = pattern
}

// Add file path to compress
func (a *ArchiveFile) AddFile(path string) {
	if a.files == nil {
		a.files = make([]string, 0)
	}
	a.files = append(a.files, path)
}

// Sets compression level
func (a *ArchiveFile) SetCompression(lvl CompressionLevel) {
	a.compression = lvl
}

// Sets volumes sizes
func (a *ArchiveFile) SetVolumes(vol string) {
	a.volumes = vol
}

// Sets password for archive
func (a *ArchiveFile) SetPassord(password string) {
	a.password = password
}

// Toggles solid flag to make is solid or not. Default - false
func (a *ArchiveFile) ToggleSolid(solid bool) {
	a.solid = solid
}

func (a *ArchiveFile) ExcludePath(extype ExcludePathFlag) {
	a.excludePath = extype
}

// Sets file encoding
func (a *ArchiveFile) SetEncoding(encoding Encoding) {
	a.encoding = encoding
}

// Compress your source to rar file to path
func (a *ArchiveFile) Compress() error {
	return a.savefile()
}

// Returns concatinated destination direactory and file name.
// If file name is empty, return "./" as current directory.
func (a *ArchiveFile) filename() string {
	return fmt.Sprintf("%s/%s.rar", a.destinationDir, a.name)
}

// Builds and returns arguments to call rar utility.
// Also returns temp file for source, to use and delete after that.
func (a *ArchiveFile) buildargs() (args []string, tempfile string, err error) {
	// first argument used for compress
	args = append(args, "a")

	if a.password != "" {
		args = append(args, fmt.Sprintf("-p%s", a.password))
	}

	if a.comment != "" {
		args = append(args, fmt.Sprintf("-z%s", a.comment))
	}

	if a.compression != NoneCompression {
		args = append(args, string(a.compression))
	}

	if a.excludePath != NotExcludePath {
		args = append(args, string(a.excludePath))
	}

	if a.encoding != DefaultEncoding {
		args = append(args, fmt.Sprintf("-sc%sl", a.encoding))
	}

	if a.volumes != "" {
		args = append(args, fmt.Sprintf("-v%s", a.volumes))
	}

	// boolean params
	if a.solid {
		args = append(args, "-s")
	}

	if a.recover {
		args = append(args, "-rr")
	}

	if a.deleteFiles {
		args = append(args, "-df")
	}

	if a.keepBroken {
		args = append(args, "-kb")
	}

	if a.timestamp {
		args = append(args, "-ts")
	}

	if a.av {
		args = append(args, "-av")
	}

	if a.ignoreAttributes {
		args = append(args, "-ac")
	}

	if a.multithreaded {
		args = append(args, "-mt")
	}

	if a.disableLock {
		args = append(args, "-ds")
	}

	// where to store
	args = append(args, a.filename())

	// source is always last param
	source, tempfile, err := a.source()
	if err != nil {
		return
	}
	args = append(args, source...)
	return
}

// Returns source of files for utility call.
// Alse creates and returns temp file fileslist*.txt to store multiple files.
func (a *ArchiveFile) source() (source []string, tempfile string, err error) {

	if a.recursive {
		source = append(source, "-r")
	}

	if a.sourceDir != "" {
		source = append(source, a.sourceDir)
	}

	if a.filePattern != "" {
		source = append(source, a.filePattern)

	}

	if a.filePattern == "" && len(a.files) > 0 {
		tempfile, err = createFilesList(a.files)
		if err != nil {
			return
		}
		source = append(source, "@"+tempfile)
	}

	return
}

// Executes command to create rar archive file
func (a *ArchiveFile) savefile() error {

	args, tempfile, err := a.buildargs()
	if err != nil {
		return err
	}

	if rarutils.ShowLogs {
		fmt.Println(rarutils.RarExeDefaultPath, strings.Join(args, " "))
	}

	_, err = cmd.Call(rarutils.RarExeDefaultPath, args)
	if err != nil {
		return err
	}

	if tempfile != "" {
		os.RemoveAll(tempfile)
	}

	return nil
}

// Creates tmp file to save as argument "@fileslist.txt" to create archive file
func createFilesList(fs []string) (string, error) {
	tempFile, err := os.CreateTemp(".", filesListFileName)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()
	fileslist := strings.Join(fs, "\n")
	tempFile.WriteString(fileslist)
	return tempFile.Name(), nil
}
