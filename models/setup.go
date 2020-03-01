package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	db.AutoMigrate(&Book{})
	db.AutoMigrate(&User{})
	return db
}
