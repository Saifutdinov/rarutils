package tests

import (
	"os"
	"testing"

	"github.com/Saifutdinov/rarutils"
	"github.com/Saifutdinov/rarutils/rar"
)

const (
	tempArchivefileName = "rar-arcive"
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

	rarutils.ShowDebugLogs(true)
	rarutils.SetRarPath("/opt/homebrew/bin/rar")

	archive := rar.NewArchiveWithConfig(rar.ArchiveConfig{
		FilePattern:    "../examples/files/example_pdf_*.pdf",
		Name:           tempArchivefileName,
		Files:          []string{tempfile1.Name(), tempfile2.Name()},
		Encoding:       rar.UTF8Encoding,
		Solid:          true,
		DestinationDir: "../examples/archives",
		ExcludePath:    rar.ExcludeBasePath,
		// Compression: rar.CompressionLVL3,
	})

	if err = archive.Compress(); err != nil {
		t.Error(err)
	}

	t.Log(tempArchivefilePath)

	// Проверяем, что архив создан
	if _, err := os.Stat("../examples/archives/rar-arcive.rar"); os.IsNotExist(err) {
		t.Errorf("Archive file was not created")
	}
}

func clear() {
	os.Remove(tempArchivefilePath)
	os.RemoveAll(tmpDestinationDir)
}
