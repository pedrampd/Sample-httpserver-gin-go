package logic

import (
	models "Desktop/go-crud/Models"
	"Desktop/go-crud/interfaces"
)

type BookLogic struct {
	bookRepository interfaces.IBookRepository
}

func NewBookLogic(bookRepo interfaces.IBookRepository) BookLogic {
	return BookLogic{bookRepository: bookRepo}
}

func (bl *BookLogic) GetAllBooks() ([]models.Book, error) {
	books, err := bl.bookRepository.GetAllBooksFromDB()
	if err != nil {
		return nil, err
	}
	return books, nil
}
