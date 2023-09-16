package logic

import (
	models "Desktop/go-crud/Models"
	"Desktop/go-crud/interfaces"
)

type AuthorLogic struct {
	authorRepository interfaces.IAuthorRepository
}

func NewAuthorLogic(authorRepo interfaces.IAuthorRepository) AuthorLogic{
	return AuthorLogic{authorRepository: authorRepo}
}

func (al *AuthorLogic) GetAllAuthors ([]models.Author, error){
	authors, err := al.authorRepository.GetAllAuthorsFromDB()
	if err != nil{
		return nil, err
	}
	return authors, nil
}