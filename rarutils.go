package rarutils

type (
	RARArchive struct {
		Name        string
		Path        string
		Compression string
		Solid       bool

		FilePattern string
		Directory   string
		Files       []string
	}
)

func main() {

}

// const (
// 	rarExe   = "/usr/local/bin/rar"
// 	unrarExe = "/usr/local/bin/unrar"

// 	actionCompress = "compress"
// 	actionExtract  = "extract"
// )

// func init() {
// 	if err := checkBinary(rarExe); err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	if err := checkBinary(unrarExe); err != nil {
// 		log.Println(err)
// 		return
// 	}
// }

// func (rar *RARArchive) AddFile(path string) {
// 	if rar.Files == nil {
// 		rar.Files = make([]string, 0)
// 	}
// 	rar.Files = append(rar.Files, path)
// }

// func (rar *RARArchive) Save() {

// }

// func (rar *RARArchive) Stream() {

// }

// // godoc
// func (rar *RARArchive) exec(action string) error {
// 	var (
// 		utility string
// 		args    = make([]string, 0)
// 	)
// 	switch action {
// 	case actionCompress:
// 		utility = rarExe
// 		args = append(args, "a")
// 		//
// 		if rar.Solid {
// 			args = append(args, "-s")
// 		}
// 		//
// 		if rar.Directory != "" || rar.FilePattern != "" {
// 			args = append(args, rar.Directory, rar.FilePattern)
// 			break
// 		}
// 		//
// 		fileslist, err := fileslist(rar.Files)
// 		if err != nil {
// 			return err
// 		}
// 		defer os.Remove(fileslist.Name())
// 		args = append(args, "@"+fileslist.Name())
// 		break
// 	case actionExtract:
// 		utility = unrarExe
// 		args = append(args, "x")
// 		break
// 	}

// 	cmd(utility, args)
// 	return nil
// }

// func fileslist(fs []string) (*os.File, error) {
// 	tempFile, err := os.CreateTemp("", "filelist.txt")
// 	if err != nil {
// 		return nil, err
// 	}
// 	fileslist := strings.Join(fs, "\n")
// 	tempFile.WriteString(fileslist)
// 	return tempFile, nil
// }
