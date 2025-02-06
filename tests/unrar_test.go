package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/Saifutdinov/rarutils"
	"github.com/Saifutdinov/rarutils/rar"
	"github.com/Saifutdinov/rarutils/unrar"
)

const tmpDestinationDir = "ArchiveName.rar_extracted"

func TestExtract(t *testing.T) {

	tempfile1, err := os.CreateTemp(".", tmpfile1)
	if err != nil {
		t.Error(err)
	}

	tempfile2, err := os.CreateTemp(".", tmpfile2)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		clear()
		os.Remove(tempfile1.Name())
		os.Remove(tempfile2.Name())
	}()

	rarutils.SetRarPath("/opt/homebrew/bin/rar")

	rarArchive := rar.NewArchive()

	rarArchive.AddFile(tempfile1.Name())
	rarArchive.AddFile(tempfile2.Name())

	rarArchive.SetCompression(rar.CompressionLVL5)
	rarArchive.ToggleSolid(true)

	if err = rarArchive.Compress(); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(tempArchivefilePath); os.IsNotExist(err) {
		t.Errorf("Archive file was not created")
	}

	rarutils.SetUnrarPath("/opt/homebrew/bin/unrar")
	unrarArchive := unrar.NewArchive("./" + tempArchivefilePath)
	os.Mkdir(tmpDestinationDir, 0755)
	unrarArchive.SetDestination(tmpDestinationDir)

	files, err := unrarArchive.Extract()
	if err != nil {
		t.Error(err)
	}

	for _, f := range files {
		if f.Name != clearname(tempfile1.Name()) && f.Name != clearname(tempfile2.Name()) {
			t.Errorf("Archive extracted wrong")
			break
		}
	}

}

func clearname(withpath string) string {
	parts := strings.Split(withpath, "/")
	return parts[len(parts)-1]
}
