package repository

import (
	"datawow/book-list/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	mySql *gorm.DB
}

func NewUserRepository(mySql *gorm.DB) *UserRepository {
	return &UserRepository{
		mySql: mySql,
	}
}

func (r *UserRepository) CreateOne(user models.User) (models.User, error) {
	if err := r.mySql.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) FindOne(id uint) (models.User, error) {
	condition := models.User{
		Model: gorm.Model{
			ID: id,
		},
	}

	user := models.User{}

	if err := r.mySql.First(&user, condition).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (models.User, error) {
	condition := models.User{
		Email: email,
	}

	user := models.User{}

	if err := r.mySql.Find(&user, condition).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
