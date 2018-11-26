package book

import "golang-tutorial/go-clean-architecture/entities"

type Service struct {
	repo BookRepository
}

func NewService (r BookRepository) *Service {
	return &Service{ repo: r }
}

func (s *Service) FetchAll() ([]*entities.Book, error) {
	return s.repo.FetchAll()
}

func (s *Service) FindById (id uint) (*entities.Book, error) {
	return s.repo.FindById(id)
}

func (s *Service) Store( b *entities.Book) (bool, error) {
	return s.repo.Store(b)
}

func (s *Service) Update( b *entities.Book) (bool, error) {
	return s.repo.Update(b)
}

func (s *Service) Delete(id uint) (bool, error) {
	return s.repo.Delete(id)
}