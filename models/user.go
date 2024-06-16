package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email    string `json:"email" gorm:"type:varchar(200);" example:"someone@example.com"`
	Name     string `json:"name" gorm:"type:varchar(200);" example:"someone"`
	Password string `json:"password" gorm:"type:varchar(200); example:"1234567890""`
}
