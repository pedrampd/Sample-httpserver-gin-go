package interfaces

import models "Desktop/go-crud/Models"

type IAuthorRepository interface {
	GetAllAuthorsFromDB() ([]models.Author, error)
}