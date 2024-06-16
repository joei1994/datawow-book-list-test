package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title     string `json:"title" gorm:"type:text" example:"Harry Potter and The Chamber of Secrets"`
	Author    string `json:"author" gorm:"type:text" example:"J.K. Rowling"`
	Genre     string `json:"genre" gorm:"type:text" example:"Fantasy"`
	Year      string `json:"year" gorm:"type:text" example:"1994"`
	Tag       string `json:"tag" gorm:"type:tag" example:"Novel"`
	CreatedBy string `gorm:"type:text" example:"someone"`
}
