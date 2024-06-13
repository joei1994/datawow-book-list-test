package domain

import "datawow/book-list/models"

type IBookRepository interface {
	CreateOne(book models.Book) (models.Book, error)
	FindOne(id uint) (models.Book, error)
	FindMany(ids []uint) ([]models.Book, error)
	DeleteOne(id uint) error
	UpdateOne(book models.Book) (models.Book, error)
}
