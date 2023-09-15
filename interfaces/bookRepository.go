package interfaces

import models "Desktop/go-crud/Models"

type IBookRepository interface {
	GetAllBooksFromDB() ([]models.Book, error)
}
