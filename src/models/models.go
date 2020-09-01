package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Animal struct {
	gorm.Model
	Name        string `gorm:"column:animalname" json:"animalname"`
	Image       string `json:"img"`
	ImagAlt     string `json:"imgalt"`
	Description string
}

func (Animal) TableName() string {
	return "animals"
}

type Customer struct {
	gorm.Model
	FirstName string  `gorm:"column:firstname" json:"firstname"`
	LastName  string  `gorm:"column:lastname" json:"lastname"`
	Email     string  `gorm:"column:email" json:"email"`
	Pass      string  `gorm:"column:password" json:"password"`
	LoggedIn  bool    `gorm:"column:loggedin" json:"loggedin"`
	Order     []Order `json:"order"`
}

func (Customer) TableName() string {
	return "customer"
}

type Order struct {
	gorm.Model
	Animal
	Customer
	CustomerID   int       `gorm:"column:customer_id" json:"customer_id"`
	AnimalID    int       `gorm:"column:animal_id" json:"animal_id"`
	OrderedDate time.Time `gorm:"column:ordered_date" json:"ordered_date"`
}

func (Order) TableName() string {
	return "order"
}
