package book

import "golang-tutorial/go-clean-architecture/entities"

type BookRepository interface {
	FetchAll() ([]*entities.Book, error)
	FindById(id uint) (*entities.Book, error)
	Store( b *entities.Book) (bool, error)
	Update( b *entities.Book) (bool, error)
	Delete(id uint) (bool, error)
}

