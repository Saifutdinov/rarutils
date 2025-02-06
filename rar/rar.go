package rar

type (
	CompressionLevel string

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

	ArchiveConfig struct {
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
		// compression. Example - m0 - m5. Default empty.
		Compression CompressionLevel
		// volumes. Example - v10MB. Default empty.
		Volumes string
		// password
		Password string
	}
)

func (cl CompressionLevel) String() string {
	return string(cl)
}

const (
	// store
	NoneCompression CompressionLevel = ""
	CompressionLVL0 CompressionLevel = "-m0"
	CompressionLVL1 CompressionLevel = "-m1"
	CompressionLVL2 CompressionLevel = "-m2"
	// default
	CompressionLVL3 CompressionLevel = "-m3"
	CompressionLVL4 CompressionLevel = "-m4"
	// maximal
	CompressionLVL5 CompressionLevel = "-m5"

	filesListFileName = "rarfileslist*"
)

var (
	DefaultArchiveConfig = ArchiveConfig{
		Name:           "rar-archive",
		DestinationDir: ".",
		Solid:          false,
		SourceDir:      ".",
		FilePattern:    "*.*",
		Files:          []string{},
		Compression:    CompressionLVL3,
		Volumes:        "",
		Password:       "",
	}
)

// Returns new Archive struct with name. Name should be without ".rar" extension
func NewArchive() *Archive {
	return NewArchiveWithConfig(DefaultArchiveConfig)
}

func NewArchiveWithConfig(config ArchiveConfig) *Archive {
	archive := new(Archive)

	if config.Name != "" {
		archive.name = config.Name
	}

	if config.DestinationDir != "" {
		archive.name = config.Name
	}

	archive.solid = config.Solid

	if config.SourceDir != "" {
		archive.sourceDir = config.SourceDir
	}

	if config.FilePattern != "" {
		archive.filePattern = config.FilePattern
	}

	if len(config.Files) > 0 {
		archive.files = config.Files
	}

	if config.Compression != NoneCompression {
		archive.compression = config.Compression
	}

	if config.Volumes != "" {
		archive.volumes = config.Volumes
	}

	if config.Password != "" {
		archive.password = config.Password
	}

	return archive
}
