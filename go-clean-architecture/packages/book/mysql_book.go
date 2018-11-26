package book

import (
	"github.com/jinzhu/gorm"
	"golang-tutorial/go-clean-architecture/entities"
)

type mysqlBookRepository struct {
	DbConnection *gorm.DB
}

func NewMysqlBookRepository(DbConn *gorm.DB) BookRepository {
	return &mysqlBookRepository{DbConnection: DbConn}
}

func (m *mysqlBookRepository) FetchAll() ([]*entities.Book, error) {
	books := make([]*entities.Book, 0)
	if err := m.DbConnection.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (m *mysqlBookRepository) FindById(id uint) (*entities.Book, error) {
	book := entities.Book{}
	if err := m.DbConnection.Find(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (m *mysqlBookRepository) Store(b *entities.Book) (bool, error) {
	if err := m.DbConnection.Create(&b).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (m *mysqlBookRepository) Update(b *entities.Book) (bool, error) {
	if err := m.DbConnection.Save(&b).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (m *mysqlBookRepository) Delete(id uint) (bool, error) {
	if err := m.DbConnection.Where("id = ?", id).Delete(entities.Book{}).Error; err != nil {
		return false, err
	}
	return true, nil
}