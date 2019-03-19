package controllers

import (
	"fmt"
	"github.com/dharmasastra/restWithIris/models"
	"github.com/kataras/iris"
	"net/http"
)

func GetAllPerson(i iris.Context) {
	var person []models.Person

	if err := db.Find(&person).Error; err != nil {
		i.StatusCode(http.StatusNotFound)
		fmt.Println(err)
	}
	_, _ = i.JSON(person)
}

func CreatePerson(i iris.Context) {
	var person models.Person
	if err := i.ReadJSON(&person); err != nil {
		i.StatusCode(http.StatusInternalServerError)
		return
	}
	db.Create(&person)
	_, _ = i.JSON(person)

}

func GetPerson(i iris.Context) {
	name := i.Params().Get("name")

	person := getPersonOr404(i, name)
	if person == nil {
		i.StatusCode(http.StatusNotFound)
		return
	}

	_, _ = i.JSON(person)
}

func UpdatePerson(i iris.Context) {
	name := i.Params().Get("name")

	person := getPersonOr404(i, name)
	if person == nil {
		i.StatusCode(http.StatusNotFound)
		return
	}
	if err := i.ReadJSON(&person); err != nil {
		i.StatusCode(http.StatusBadRequest)
		return
	}
	if err := db.Save(&person).Error; err != nil {
		i.StatusCode(http.StatusInternalServerError)
		return
	}

	_, _ = i.JSON(person)

}

func getPersonOr404(i iris.Context, name string) *models.Person {
	person := models.Person{}

	if err := db.First(&person, models.Person{FirstName: name}).Error; err != nil {
		i.StatusCode(http.StatusNotFound)
		fmt.Println(err)
		return nil
	}
	return &person
}
