package models

import "gorm.io/gorm"

type Role string;

const (
	RoleWaiter Role = "waiter"
	RoleChef Role = "chef"
	RoleCashier Role = "cashier"
	RoleManager Role = "manager"
	RoleAdmin Role = "admin"
)

type User struct {
	gorm.Model
	ID	   uint   `json:"id" gorm:"primaryKey"`;
	FirstName string `json:"first_name"`;
	LastName string `json:"last_name"`;
	Email string `json:"email" gorm:"unique"`;
	Role Role `json:"role"`
	Password string `json:"-" gorm:"not null"`;
}
