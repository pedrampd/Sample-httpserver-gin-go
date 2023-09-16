package api

import (
	"Desktop/go-crud/interfaces"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controllerAuthor struct {
	authorLogic interfaces.IAuthorService
}

func NewControllerAuthor(al interfaces.IAuthorService) controllerAuthor {
	return controllerAuthor{authorLogic: al}
}

func (controller *controllerAuthor) GetAuthors() (c *gin.Context) {
	books, err := controller.authorLogic.GetAllAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	byteData, err := json.Marshal(books)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, byteData)
}
