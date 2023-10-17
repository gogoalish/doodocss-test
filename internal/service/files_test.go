package service

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompressFiles(t *testing.T) {
	filesService := &Files{}
	// files := []*os.File{}
	filesData := [][]byte{}
	var filePaths, fileNames []string
	for i := 0; i < 2; i++ {
		text := fmt.Sprintf("test%d.txt", i)
		f, err := os.Create(text)
		require.NoError(t, err)
		defer os.Remove(f.Name())
		f.WriteString(text)
		filesData = append(filesData, []byte(text))
		fileNames = append(fileNames, f.Name())
		filePaths = append(filePaths, f.Name())
	}

	zipData, err := filesService.CompressFiles(filePaths, fileNames)

	require.NoError(t, err)
	require.NotNil(t, zipData)

	reader, err := zip.NewReader(bytes.NewReader(*zipData), int64(len(*zipData)))
	require.NoError(t, err)
	for i, file := range reader.File {
		data, err := file.Open()
		require.NoError(t, err)

		fileContent, err := io.ReadAll(data)
		require.NoError(t, err)

		require.Equal(t, string(filesData[i]), string(fileContent))
	}
}
