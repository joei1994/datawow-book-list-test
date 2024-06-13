package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title  string `json:"title" gorm:"type:text"`
	Author string `json:"author" gorm:"type:text"`
	Genre  string `json:"genre" gorm:"type:text"`
	Year   string `json:"year" gorm:"type:text"`
}
