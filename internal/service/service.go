package service

type Service struct {
	ArchiveService
}

func NewService() *Service {
	return &Service{
		ArchiveService: NewArchive(),
	}
}
