package unrar

import "os"

type (
	Archive struct {
		Name       string
		SourceFile string
		Password   string

		NotOverwrite bool
		Destination  string

		NotClear bool

		Files []*os.File

		ExeFilePath string
	}

	Fileinfo struct {
		Name string
		Size int64
	}
)

const (
	ActionExtract = "x"
	ActionList    = "l"
)

var (
	UnrarExeFile = "/usr/local/bin/unrar"
)

func SetEXEPath(path string) {
	UnrarExeFile = path
}
