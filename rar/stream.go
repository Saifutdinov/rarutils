package rar

import (
	"bytes"
	"io"
	"os"
	"time"
)

type (
	ArchiveStream struct {
		Name    string
		Size    int64
		ModTime time.Time
		Content io.ReadSeeker
	}
)

// Creates file with params, returns you *ArchiveStream
// and then removes file from path if param `keepAfterReturn` is false (default: false)
func (a *ArchiveFile) Stream(keepAfterReturn ...bool) (*ArchiveStream, error) {
	if err := a.savefile(); err != nil {
		return nil, err
	}

	tmpfile, err := os.Open(a.filename())
	if err != nil {
		return nil, err
	}
	defer tmpfile.Close()

	if len(keepAfterReturn) > 0 && !keepAfterReturn[0] {
		defer func(name string) {
			_ = os.Remove(name)
		}(a.filename())
	}

	info, err := tmpfile.Stat()
	if err != nil {
		return nil, err
	}

	archivefile, err := io.ReadAll(tmpfile)
	if err != nil {
		return nil, err
	}

	return &ArchiveStream{
		Name:    a.name,
		Size:    info.Size(),
		ModTime: info.ModTime(),
		Content: bytes.NewReader(archivefile),
	}, nil
}
