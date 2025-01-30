package rar

type (
	CompressionLevel int

	Archive struct {
		// how to name file. No ".rar" needed in the end.
		name string
		// where save file
		destinationDir string
		// save as solid
		solid bool
		//Directory of files. Example - /path/to/directory
		sourceDir string
		//File pattern of files. Example - /path/to/files/*.pdf
		filePattern string
		// List of file paths. Example - [/path/to/file1.pdf, /path/to/file2.pdf, /path/to/file3.pdf, ...]
		files []string
		// compression. Example - m0 - m5. Default empty.
		compression CompressionLevel
		// volumes. Example - v10MB. Default empty.
		volumes string
		// password
		password string
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

	filesListFileName = "rarfileslist*"
)

// Returns new Archive struct with name. Name should be without ".rar" extension
func NewArchive(name string) *Archive {
	return &Archive{
		name: name,
	}
}
