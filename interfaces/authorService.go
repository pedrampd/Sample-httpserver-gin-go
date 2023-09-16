package interfaces

import models "Desktop/go-crud/Models"

type IAuthorService interface {
	GetAllAuthors() ([]models.Author, error)
}