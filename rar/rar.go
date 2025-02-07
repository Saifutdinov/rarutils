package rar

type (
	Encoding         string
	CompressionLevel string
	ExcludePathFlag  string

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
		// exludes or includes path of files
		excludePath ExcludePathFlag
		// charater set for filenames -scUTF-8
		encoding Encoding
	}

	ArchiveConfig struct {
		Name           string
		DestinationDir string
		Solid          bool
		SourceDir      string
		FilePattern    string
		Files          []string
		Compression    CompressionLevel
		Volumes        string
		Password       string
		ExcludePath    ExcludePathFlag
		Encoding       Encoding
	}
)

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
)

const (
	NotExcludePath  ExcludePathFlag = ""
	ExcludePath     ExcludePathFlag = "-ep"
	ExcludeBasePath ExcludePathFlag = "-ep1"
	ExcludePathFull ExcludePathFlag = "-ep3"
)

const (
	UTF8 Encoding = "UTF-8"
)

const filesListFileName = "rarfileslist*"

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
		ExcludePath:    NotExcludePath,
		Encoding:       "",
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
		archive.destinationDir = config.DestinationDir
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

	if config.ExcludePath != NotExcludePath {
		archive.excludePath = config.ExcludePath
	}

	return archive
}
