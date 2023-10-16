package service

type Service struct {
	ArchiveService
	FilesService
}

func NewService() *Service {
	return &Service{
		ArchiveService: &Archive{},
		FilesService:   &Files{},
	}
}
