package repository

import "datawow/book-list/infrastructure"

type Repository struct {
	BookRepository *BookRepository
	UserRepository *UserRepository
}

func NewRepository(infra infrastructure.Infrastructure) Repository {
	return Repository{
		BookRepository: NewBookRepository(infra.MySql),
		UserRepository: NewUserRepository(infra.MySql),
	}
}
