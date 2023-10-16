package service

import (
	"archive/zip"
	"os"
	"testing"
)

func TestArchive_GetInformation(t *testing.T) {
	// Подготовка тестовых данных: создание временного архива
	testArchiveName := "test_archive.zip"
	testArchiveContent := []byte("Test data")

	err := createTestArchive(testArchiveName, testArchiveContent)
	if err != nil {
		t.Fatalf("Failed to create test archive: %v", err)
	}
	defer os.Remove(testArchiveName)

	archive := &Archive{}
	expectedArcInfo := &ArchiveInfo{
		FileName:    "test_archive.zip",
		ArchiveSize: float64(len(testArchiveContent)),
		TotalFiles:  1,
		TotalSize:   float64(len(testArchiveContent)),
		Files: []*FileInfo{
			{
				FilePath: "test_archive.zip",
				Size:     float64(len(testArchiveContent)),
				MimeType: "application/zip",
			},
		},
	}

	// Вызов тестируемой функции
	arcInfo, err := archive.GetInformation(testArchiveName, "test_archive.zip")

	// Проверка результатов
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if arcInfo.FileName != expectedArcInfo.FileName {
		t.Errorf("Expected FileName to be %s, got %s", expectedArcInfo.FileName, arcInfo.FileName)
	}

	if arcInfo.ArchiveSize != expectedArcInfo.ArchiveSize {
		t.Errorf("Expected ArchiveSize to be %v, got %v", expectedArcInfo.ArchiveSize, arcInfo.ArchiveSize)
	}

	if arcInfo.TotalFiles != expectedArcInfo.TotalFiles {
		t.Errorf("Expected TotalFiles to be %d, got %d", expectedArcInfo.TotalFiles, arcInfo.TotalFiles)
	}

	if arcInfo.TotalSize != expectedArcInfo.TotalSize {
		t.Errorf("Expected TotalSize to be %f, got %f", expectedArcInfo.TotalSize, arcInfo.TotalSize)
	}

	if len(arcInfo.Files) != len(expectedArcInfo.Files) {
		t.Errorf("Expected %d files, got %d", len(expectedArcInfo.Files), len(arcInfo.Files))
	}

	for i, expectedFile := range expectedArcInfo.Files {
		if arcInfo.Files[i].FilePath != expectedFile.FilePath {
			t.Errorf("Expected FilePath to be %s, got %s", expectedFile.FilePath, arcInfo.Files[i].FilePath)
		}

		if arcInfo.Files[i].Size != expectedFile.Size {
			t.Errorf("Expected Size to be %f, got %f", expectedFile.Size, arcInfo.Files[i].Size)
		}

		if arcInfo.Files[i].MimeType != expectedFile.MimeType {
			t.Errorf("Expected MimeType to be %s, got %s", expectedFile.MimeType, arcInfo.Files[i].MimeType)
		}
	}
}

func createTestArchive(archiveName string, content []byte) error {
	file, err := os.Create(archiveName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Создаем zip-архив с одним файлом
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	fileWriter, err := zipWriter.Create(archiveName)
	if err != nil {
		return err
	}

	_, err = fileWriter.Write(content)
	if err != nil {
		return err
	}

	return nil
}
