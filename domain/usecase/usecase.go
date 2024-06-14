package domain

import (
	"datawow/book-list/config"
	"datawow/book-list/repository"
)

type UseCase struct {
	BookUseCase  *BookUseCase
	UserUseCase  *UserUseCase
	TokenUseCase *TokenUseCase
}

func NewUseCase(config *config.Config, repository repository.Repository) UseCase {
	tokenUseCase := NewTokenUseCase(config)

	return UseCase{
		BookUseCase:  NewBookUseCase(repository.BookRepository),
		UserUseCase:  NewUserUseCase(repository.UserRepository, tokenUseCase),
		TokenUseCase: tokenUseCase,
	}
}
