package rar

// import "../rarutils"

type (
	Archive struct {
		// how to name file
		Name string
		// where save file
		Path string

		// save as solid
		Solid bool

		// source

		//Directory of files. Example - /path/to/directory
		Directory string
		//File pattern of files. Example - /path/to/files/*.pdf
		FilePattern string
		// List of file paths. Example - [/path/to/file1.pdf, /path/to/file2.pdf, /path/to/file3.pdf, ...]
		Files []string
	}
)

func init() {
	// rarutils.CheckBinary()
}

func args() {

}

func cmd() {

}

func (a Archive) Save() {

}
