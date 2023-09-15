package interfaces

import models "Desktop/go-crud/Models"

type IbookService interface {
	GetAllBooks() ([]models.Book, error)
}
