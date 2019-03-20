package models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	FirstName string    `gorm:"unique" json:"first_name"`
	LastName  string    `gorm:"unique" json:"last_name"`
	Addresses []Address `gorm:"ForeignKey:IdAddress" json:"addresses"`
}

type Address struct {
	IdAddress uint
	Street    string `json:"street"`
	City      string `json:"city"`
}
