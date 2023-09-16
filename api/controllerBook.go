package api

import (
	"Desktop/go-crud/interfaces"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerBook struct {
	bookService interfaces.IbookService
}

func NewControllerBook(bs interfaces.IbookService) ControllerBook {
	return ControllerBook{bookService: bs}
}

func (controller *ControllerBook) GetBooks(c *gin.Context) {
	books, err := controller.bookService.GetAllBooks()
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
