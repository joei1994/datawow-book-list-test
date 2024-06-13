package domain

import "datawow/book-list/repository"

type UseCase struct {
	BookUseCase *BookUseCase
}

func NewUseCase(repository repository.Repository) UseCase {
	return UseCase{
		BookUseCase: NewBookUseCase(repository.BookRepository),
	}
}
