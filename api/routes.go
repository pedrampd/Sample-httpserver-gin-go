package api

import (
	"Desktop/go-crud/logic"
	"Desktop/go-crud/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetBookRoutes(server *gin.Engine, db *gorm.DB) {
	bookRepository := repository.NewBookRepository(db)
	bookService := logic.NewBookLogic(&bookRepository)
	bookController := NewControllerBook(&bookService)

	groupBook := server.Group("/book")
	groupBook.GET("/all", bookController.GetBooks)
}
