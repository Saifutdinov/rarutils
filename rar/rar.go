package rar

type (
	CompressionLevel int

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
		Compression CompressionLevel
		// Volumes. Example - v10MB. Default empty.
		Volumes string
		// Password
		Password string
	}
)

const (
	// store
	CompressionLVL0 CompressionLevel = iota
	CompressionLVL1
	CompressionLVL2
	// default
	CompressionLVL3
	CompressionLVL4
	// maximal
	CompressionLVL5

	FilesListFileName = "rarfileslist*"
)

var (
	RarExeFile = "/usr/local/bin/rar"
)

// Returns new Archive struct with name. Name should be without ".rar" extension
func NewArchive(name string) *Archive {
	return &Archive{
		Name: name,
	}
}

// Sets path to executable file
func SetEXEPath(path string) {
	RarExeFile = path
}
