package unrar

type (
	Archive struct {
		Name       string
		SourceFile string
		Password   string

		NotOverwrite bool
		Destination  string
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

func NewArchive(path string) *Archive {
	return &Archive{
		SourceFile: path,
	}
}
