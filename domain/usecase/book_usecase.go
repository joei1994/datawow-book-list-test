package domain

import (
	repository "datawow/book-list/domain/repository"
	"datawow/book-list/models"
)

type IBookUseCase interface {
	CreateBook(book models.Book) (models.Book, error)
	GetBook(id uint) (models.Book, error)
	GetBooks(ids []uint) ([]models.Book, error)
	DeleteBook(id uint) error
	UpdateBook(id uint, book models.Book) (models.Book, error)
}

type BookUseCase struct {
	bookRepo repository.IBookRepository
}

func NewBookUseCase(bookRepo repository.IBookRepository) *BookUseCase {
	return &BookUseCase{
		bookRepo: bookRepo,
	}
}

func (uc *BookUseCase) CreateBook(book models.Book) (models.Book, error) {
	return uc.bookRepo.CreateOne(book)
}

func (uc *BookUseCase) GetBook(id uint) (models.Book, error) {
	return uc.bookRepo.FindOne(id)
}

func (uc *BookUseCase) GetBooks(ids []uint) ([]models.Book, error) {
	return uc.bookRepo.FindMany(ids)
}

func (uc *BookUseCase) DeleteBook(id uint) error {
	return uc.bookRepo.DeleteOne(id)
}

func (uc *BookUseCase) UpdateBook(id uint, book models.Book) (models.Book, error) {
	book.ID = id
	return uc.bookRepo.UpdateOne(book)
}
