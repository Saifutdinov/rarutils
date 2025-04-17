package rar

type (
	Encoding         string
	CompressionLevel string
	ExcludePathFlag  string

	ArchiveFile struct {
		// how to name file. No ".rar" needed in the end.
		name string
		// where save file
		destinationDir string
		//Directory of files. Example - /path/to/directory
		sourceDir string
		//File pattern of files. Example - /path/to/files/*.pdf
		filePattern string
		// volumes. Example - v10MB.
		volumes string
		// password
		password string
		// -z<file>
		comment string
		// List of file paths. Example - [/path/to/file1.pdf, /path/to/file2.pdf, /path/to/file3.pdf, ...]
		files []string
		// compression. Example - m0 - m5.
		compression CompressionLevel
		// exludes or includes path of files
		excludePath ExcludePathFlag
		// charater set for filenames -scUTF-8
		encoding Encoding
		// save as solid
		solid bool
		// save recursive all directories
		recursive bool
		// stores recovery file (3%)
		recover bool
		// delete files after all
		deleteFiles bool
		// keep broken files
		keepBroken bool
		//
		timestamp bool
		//
		av bool

		ignoreAttributes bool

		multithreaded bool

		disableLock bool
	}

	ArchiveConfig struct {
		// archivename.rar (default: "rar-archive")
		Name string
		// /path/to/store/archive (default:".")
		DestinationDir string
		// /path/to/compress (default:".")
		SourceDir string
		// file*.pdf (default: "*.*")
		FilePattern string
		// -v<size> (default: "10m")
		VolumeSize string
		// -p[secretpass]
		Password string
		// -z<file> - text from <file> as comment
		CommentFile string
		// /path/to/file
		Files []string
		// -sc[encoding] (defualt: "")
		Encoding Encoding
		// -m[0-5] (default: CompressionLVL3)
		Compression CompressionLevel
		// -ep (default: false)
		ExcludePath ExcludePathFlag
		// -s: solid archive (default: false)
		Solid bool
		// -r (default: false)
		Recursive bool
		// -rr (defualt: false)
		RecoveryRecord bool
		// -df (defualt: false) - USE CAREFULLY!
		DeleteFiles bool
		// -kb (defualt: false)
		KeepBroken bool
		// -ts (defualt: true)
		TimeStamp bool
		// -av checked with antivirus (defualt: false)
		AV bool
		// -ac ignore files attributes (like "archived") for reserve
		IgnoreAttributes bool
		// -mt Use thtreads to compress.
		Multithreaded bool
		// -ds Disable reading lock (USE CAREFULLY!)
		DisableLock bool
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
	DefaultEncoding Encoding = ""
	UTF16Encoding   Encoding = "u"
	UTF8Encoding    Encoding = "f"
	ANSIEncoding    Encoding = "a" // windows only
	DOSEncoding     Encoding = "o" // windows only

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
		VolumeSize:     "10m",
		Password:       "",
		ExcludePath:    NotExcludePath,
		Encoding:       DefaultEncoding,
	}
)

// Returns new Archive struct with name. Name should be without ".rar" extension
func NewArchive() *ArchiveFile {
	return NewArchiveWithConfig(DefaultArchiveConfig)
}

func NewArchiveWithConfig(config ArchiveConfig) *ArchiveFile {
	archive := new(ArchiveFile)

	// string params
	if config.Name != "" {
		archive.name = config.Name
	}

	if config.DestinationDir != "" {
		archive.destinationDir = config.DestinationDir
	}

	if config.SourceDir != "" {
		archive.sourceDir = config.SourceDir
	}

	if config.FilePattern != "" {
		archive.filePattern = config.FilePattern
	}

	if config.VolumeSize != "" {
		archive.volumes = config.VolumeSize
	}

	if config.Password != "" {
		archive.password = config.Password
	}

	if config.CommentFile != "" {
		archive.comment = config.CommentFile
	}

	// const params
	if config.ExcludePath != NotExcludePath {
		archive.excludePath = config.ExcludePath
	}

	if config.Encoding != DefaultEncoding {
		archive.encoding = config.Encoding
	}

	if config.Compression != NoneCompression {
		archive.compression = config.Compression
	}

	// slice
	if len(config.Files) > 0 {
		archive.files = config.Files
	}

	// bool
	archive.solid = config.Solid
	archive.recursive = config.Recursive
	archive.recover = config.RecoveryRecord
	archive.deleteFiles = config.DeleteFiles
	archive.keepBroken = config.KeepBroken
	archive.timestamp = config.TimeStamp
	archive.av = config.AV
	archive.ignoreAttributes = config.IgnoreAttributes
	archive.multithreaded = config.Multithreaded
	archive.disableLock = config.DisableLock

	return archive
}
