package domain

import "datawow/book-list/models"

type IUserRepository interface {
	CreateOne(user models.User) (models.User, error)
	FindOne(id uint) (models.User, error)
	FindByEmail(email string) (models.User, error)
}
