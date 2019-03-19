package controllers

import (
	"github.com/dharmasastra/restWithIris/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	conn, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}

	db = conn
	db.AutoMigrate(&models.Person{})
}
