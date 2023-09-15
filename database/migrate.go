package database

import (
	models "Desktop/go-crud/Models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.User{})
}
