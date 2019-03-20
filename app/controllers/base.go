package controllers

import (
	"fmt"
	"github.com/dharmasastra/restWithIris/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbConnection := os.Getenv("DB_CONNECTION")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)

	conn, err := gorm.Open(dbConnection, dbUri)
	if err != nil {
		panic(err)
	}

	db = conn
	db.AutoMigrate(&models.Person{}, &models.Address{})
	db.LogMode(true)
}
