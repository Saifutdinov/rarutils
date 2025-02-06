package tests

import (
	"os"
	"testing"

	"github.com/Saifutdinov/rarutils"
	"github.com/Saifutdinov/rarutils/rar"
)

const (
	tempArchivefileName = "ArchiveName"
	tempArchivefilePath = tempArchivefileName + ".rar"
	tmpfile1            = "tempfile1.pdf"
	tmpfile2            = "tempfile2.pdf"
)

func TestCompress(t *testing.T) {
	defer clear()

	tempfile1, err := os.CreateTemp("", tmpfile1)
	if err != nil {
		t.Error(err)
	}

	tempfile2, err := os.CreateTemp("", tmpfile2)
	if err != nil {
		t.Error(err)
	}

	rarutils.SetRarPath("/opt/homebrew/bin/rar")

	archive := rar.NewArchive()

	archive.AddFile(tempfile1.Name())
	archive.AddFile(tempfile2.Name())

	archive.SetCompression(rar.CompressionLVL5)
	archive.ToggleSolid(true)

	if err = archive.Compress(); err != nil {
		t.Error(err)
	}

	// Проверяем, что архив создан
	if _, err := os.Stat(tempArchivefilePath); os.IsNotExist(err) {
		t.Errorf("Archive file was not created")
	}
}

func clear() {
	os.Remove(tempArchivefilePath)
	os.RemoveAll(tmpDestinationDir)
}
