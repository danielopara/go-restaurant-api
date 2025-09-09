package models

import "gorm.io/gorm"

type Category string

const (
	Food    Category = "food"
	Snack   Category = "snack"
	Drink   Category = "drink"
	Platter Category = "platter"
)

type Menu struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`;
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Category Category `json:"category"`
	Available bool `json:"available"`
	CreatedBy string `json:"created_by"`
}