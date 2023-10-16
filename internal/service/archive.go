package service

import (
	"archive/zip"

	"github.com/gogoalish/doodocs-test/utils"
)

type ArchiveService interface {
	GetInformation(src, arcName string) (*ArchiveInfo, error)
}

type Archive struct{}

func NewArchive() *Archive {
	return &Archive{}
}

type ArchiveInfo struct {
	FileName    string      `json:"filename"`
	ArchiveSize float64     `json:"archive_size"`
	TotalSize   float64     `json:"total_size"`
	TotalFiles  int         `json:"total_files"`
	Files       []*FileInfo `json:"files"`
}

type FileInfo struct {
	FilePath string  `json:"file_path"`
	Size     float64 `json:"size"`
	MimeType string  `json:"mimetype"`
}

func (archive *Archive) GetInformation(src, arcName string) (*ArchiveInfo, error) {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	arcSize, err := utils.GetFileSize(src)
	if err != nil {
		return nil, err
	}

	arcInfo := &ArchiveInfo{
		FileName:    arcName,
		ArchiveSize: arcSize,
		TotalFiles:  len(reader.File),
	}
	for _, file := range reader.File {
		fileSize := file.UncompressedSize64
		fileInfo := &FileInfo{
			FilePath: file.Name,
			Size:     float64(fileSize),
			MimeType: utils.DetectMimeType(file.Name),
		}
		arcInfo.TotalSize += fileInfo.Size
		arcInfo.Files = append(arcInfo.Files, fileInfo)
	}
	return arcInfo, nil
}
