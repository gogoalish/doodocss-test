package service

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInformation(t *testing.T) {
	archiveService := &Archive{}

	src := "testarchive.zip"
	arcName := "testarchive"

	err := createTestArchive(src, t)
	require.NoError(t, err)
	defer os.Remove(src)

	arcInfo, err := archiveService.GetInformation(src, arcName)

	require.NoError(t, err)
	require.NotNil(t, arcInfo)

	require.Equal(t, arcName, arcInfo.FileName)
	require.NotEqual(t, 0.0, arcInfo.ArchiveSize)
	require.NotEqual(t, 0.0, arcInfo.TotalSize)
	require.True(t, arcInfo.TotalFiles > 0)

}

func createTestArchive(src string, t *testing.T) error {
	zipFile, err := os.Create(src)
	require.NoError(t, err)

	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	for i := 0; i < 2; i++ {
		fileName := fmt.Sprintf("%d.txt", i)
		file, _ := os.Open(fileName)
		defer os.Remove(fileName)
		file.WriteString("test")
		fileWriter, err := zipWriter.Create(fileName)
		require.NoError(t, err)

		_, err = io.Copy(fileWriter, file)
		require.NoError(t, err)
	}
	return nil
}
