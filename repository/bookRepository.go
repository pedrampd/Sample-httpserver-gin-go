package repository

import (
	models "Desktop/go-crud/Models"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return BookRepository{db: db}
}

func (br *BookRepository) GetAllBooksFromDB() ([]models.Book, error) {
	var books []models.Book
	dbCtx := br.db.Find(&books)
	if dbCtx.Error != nil {
		return nil, dbCtx.Error
	}
	return books, nil
}
