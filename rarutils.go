package rarutils

import "github.com/Saifutdinov/rarutils/rar"

func main() {

	var (
		fileName    = "allfiles"
		source      = "./files"
		destination = "./archives"
	)

	archive := rar.Archive{
		Name:           fileName,
		SourceDir:      source,
		DestinationDir: destination,
	}

	archive.Save()

}
