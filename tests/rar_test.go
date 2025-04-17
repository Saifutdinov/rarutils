package tests

import (
	"os"
	"testing"

	"github.com/Saifutdinov/rarutils"
	"github.com/Saifutdinov/rarutils/rar"
)

const (
	exampleArchivesPath = "../examples/archives"
	exampleFilesPath    = "../examples/files"
	tempArchivefileName = "rar-archive"
	tempArchivefilePath = exampleArchivesPath + "/" + tempArchivefileName + ".rar"
)

func TestCompress(t *testing.T) {
	// defer clear()

	rarutils.ShowDebugLogs(true)
	rarutils.SetRarPath("/opt/homebrew/bin/rar")

	archive := rar.NewArchiveWithConfig(rar.ArchiveConfig{
		FilePattern:    "../examples/files/example_pdf_*.pdf",
		Encoding:       rar.UTF8Encoding,
		Solid:          true,
		DestinationDir: exampleArchivesPath,
		ExcludePath:    rar.ExcludeBasePath,
		// Compression: rar.CompressionLVL3,
	})

	if err := archive.Compress(); err != nil {
		t.Error(err)
		return
	}

	t.Log(tempArchivefilePath)

	// Проверяем, что архив создан
	if _, err := os.Stat(tempArchivefilePath); os.IsNotExist(err) {
		t.Errorf("Archive file was not created")
	}
}

func clear() {
	os.Remove(tempArchivefilePath)
}
