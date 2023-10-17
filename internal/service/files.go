package service

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
)

type FilesService interface {
	CompressFiles(filePaths, fileNames []string) (*[]byte, error)
}

type Files struct{}

func (*Files) CompressFiles(filePaths, fileNames []string) (*[]byte, error) {
	zipBuffer := new(bytes.Buffer)
	writer := zip.NewWriter(zipBuffer)

	for i, path := range filePaths {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		fileWriter, err := writer.Create(fileNames[i])
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(fileWriter, file)
		if err != nil && err != io.EOF {
			return nil, err
		}
	}

	writer.Close()
	zipData := zipBuffer.Bytes()
	return &zipData, nil
}
