package repository

import (
	models "Desktop/go-crud/Models"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository{
	return AuthorRepository{db : db}
}

func (ar *AuthorRepository) GetAllAuthorsFromDB() ([]models.Author, error){
	var authors []models.Author
	dbCtx := ar.db.Find(&authors)
	if dbCtx.Error != nil{
		return nil, dbCtx.Error
	}
	return authors, nil
}