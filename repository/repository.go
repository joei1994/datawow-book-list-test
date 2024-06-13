package repository

import "datawow/book-list/infrastructure"

type Repository struct {
	BookRepository *BookRepository
}

func NewRepository(infra infrastructure.Infrastructure) Repository {
	return Repository{
		BookRepository: NewBookRepository(infra.MySql),
	}
}
