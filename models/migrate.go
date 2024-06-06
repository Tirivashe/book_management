package models

import (
	"github.com/Tirivashe/book_management/db"
	"gorm.io/gorm"
)

var database *gorm.DB
func init() {
	db.ConnectDatabase()
	database = db.GetDB()
	database.AutoMigrate(&Book{})
}
