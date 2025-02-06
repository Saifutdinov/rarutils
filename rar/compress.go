package rar

import (
	"fmt"
	"os"
	"strings"

	"github.com/Saifutdinov/rarutils"
	"github.com/Saifutdinov/rarutils/cmd"
)

// Sets source directory to compress
func (a *Archive) SetSourceDir(dir string) {
	a.sourceDir = dir
}

// Sets destination directory to store archive
func (a *Archive) SetDestinationDir(dir string) {
	a.destinationDir = dir
}

// Sets pattern of file to compress
func (a *Archive) SetFilePattern(pattern string) {
	a.filePattern = pattern
}

// Add file path to compress
func (a *Archive) AddFile(path string) {
	if a.files == nil {
		a.files = make([]string, 0)
	}
	a.files = append(a.files, path)
}

// Sets compression level
func (a *Archive) SetCompression(lvl CompressionLevel) {
	a.compression = CompressionLevel(lvl)
}

// Sets volumes sizes
func (a *Archive) SetVolumes(vol string) {
	a.volumes = vol
}

// Sets password for archive
func (a *Archive) SetPassord(password string) {
	a.password = password
}

// Toggles solid flag to make is solid or not. Default - false
func (a *Archive) ToggleSolid(solid bool) {
	a.solid = solid
}

// Compress your source to rar file to path
func (a *Archive) Compress() error {
	return a.savefile()
}

// Creates file with params, returns you []byte to force download or send by email
// and then removes file from path (you need just []byte)
func (a *Archive) Stream(keepAfterReturn bool) (*os.File, error) {
	if err := a.savefile(); err != nil {
		return nil, err
	}
	file, err := os.Open(a.filename())
	if err != nil {
		return nil, err
	}
	if !keepAfterReturn {
		os.Remove(a.filename())
	}
	return file, nil
}

// Returns concatinated destination direactory and file name.
// If file name is empty, return "./" as current directory.
func (a *Archive) filename() string {
	if a.destinationDir == "" {
		a.destinationDir = "."
	}
	return fmt.Sprintf("%s/%s.rar", a.destinationDir, a.name)
}

// Builds and returns arguments to call rar utility.
// Also returns temp file for source, to use and delete after that.
func (a *Archive) buildargs() (args []string, tempfile string, err error) {
	args = append(args, "a", a.filename())
	if a.solid {
		args = append(args, "-s")
	}

	if a.compression != CompressionLVL3 {
		args = append(args, fmt.Sprintf("-m%d", a.compression))
	}

	if a.volumes != "" {
		args = append(args, "-v"+a.volumes)
	}
	source, tempfile, err := a.source()
	if err != nil {
		return
	}
	args = append(args, source...)
	if a.password != "" {
		args = append(args, "-p"+a.password)
	}
	return
}

// Returns source of files for utility call.
// Alse creates and returns temp file fileslist*.txt to store multiple files.
func (a *Archive) source() (source []string, tempfile string, err error) {
	if a.sourceDir != "" {
		source = append(source, "-r", a.sourceDir)
		return
	}
	if a.filePattern != "" {
		source = append(source, a.filePattern)
		return
	}

	tempfile, err = createFilesList(a.files)
	if err != nil {
		return
	}
	source = append(source, "@"+tempfile)
	return
}

// Executes command to create rar archive file
func (a *Archive) savefile() error {
	args, tempfile, err := a.buildargs()
	if err != nil {
		return err
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
	tempFile, err := os.CreateTemp("", filesListFileName)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()
	fileslist := strings.Join(fs, "\n")
	tempFile.WriteString(fileslist)
	return tempFile.Name(), nil
}
