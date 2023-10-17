package service

type Service struct {
	ArchiveService
	FilesService
	MailService
}

func NewService() *Service {
	return &Service{
		ArchiveService: &Archive{},
		FilesService:   &Files{},
		MailService:    &Mail{},
	}
}
