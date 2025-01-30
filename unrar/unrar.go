package unrar

type (
	Archive struct {
		// Path to file with file name /path/to/file.rar.
		sourceFile string
		// Password to extract protected archive.
		password string

		// Overwrite extracted files. False by default.
		notOverwrite bool

		// Path to extract files, default ./{file name}_extracted_{time.Now().Nanoseconds()}
		destination string
	}

	Fileinfo struct {
		// parsed file name
		Name string
		// parsed file size
		Size int64
	}
)

const (
	actionExtract = "x"
	actionList    = "l"
)

func NewArchive(path string) *Archive {
	return &Archive{
		sourceFile: path,
	}
}
