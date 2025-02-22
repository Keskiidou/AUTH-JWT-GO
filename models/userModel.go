package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"unique"`
	Password string
	Phone    int `gorm:"unique"`
	Address  string
}
