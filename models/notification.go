package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	UserID uint	`json:"units_id"`
	Role string `json:"role"`
	Type string  `json:"type"`
	Message string	`json:"message`
}