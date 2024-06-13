package repository

import (
	"datawow/book-list/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	mySql *gorm.DB
}

func NewBookRepository(mySql *gorm.DB) *BookRepository {
	return &BookRepository{
		mySql: mySql,
	}
}

func (r *BookRepository) CreateOne(book models.Book) (models.Book, error) {
	if err := r.mySql.Create(&book).Error; err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (r *BookRepository) FindOne(id uint) (models.Book, error) {
	condition := models.Book{
		Model: gorm.Model{
			ID: id,
		},
	}

	book := models.Book{}

	if err := r.mySql.First(&book, condition).Error; err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (r *BookRepository) FindMany(ids []uint) ([]models.Book, error) {
	books := []models.Book{}

	if err := r.mySql.Find(&books, "id IN ?", ids).Error; err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func (r *BookRepository) DeleteOne(id uint) error {
	condition := models.Book{
		Model: gorm.Model{
			ID: id,
		},
	}

	book := models.Book{}

	if err := r.mySql.Delete(&book, condition).Error; err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) UpdateOne(book models.Book) (models.Book, error) {
	if err := r.mySql.Updates(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}
