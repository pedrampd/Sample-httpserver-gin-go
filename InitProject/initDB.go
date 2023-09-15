package initproject

import (
	"Desktop/go-crud/database"
	"log"

	// "gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("C:\\Users\\Pedram\\Desktop\\go-crud\\sqlite\\database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Could'nt connect to db")
		return nil, err
	}
	database.Migrate(db)
	return db, nil

}
