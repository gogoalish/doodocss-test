package utils

import (
	"mime"
	"os"
	"path/filepath"
)

func DetectMimeType(path string) string {
	return mime.TypeByExtension(filepath.Ext(path))
}

func GetFileSize(fileName string) (float64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return float64(fileInfo.Size()), nil
}

func IsAllowedType(filename string, allowedTypes []string) bool {
	filetype := DetectMimeType(filename)
	for _, allowed := range allowedTypes {
		if allowed == filetype {
			return true
		}
	}
	return false
}
