package domain

import (
	"errors"

	repository "datawow/book-list/domain/repository"
	"datawow/book-list/models"

	"golang.org/x/crypto/bcrypt"
)

type IUserUseCase interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(id uint) (models.User, error)
	Register(user models.User) (models.User, error)
}

type UserUseCase struct {
	userRepo repository.IUserRepository
}

func NewUserUseCase(userRepo repository.IUserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) CreateUser(user models.User) (models.User, error) {
	return uc.userRepo.CreateOne(user)
}

func (uc *UserUseCase) GetUser(id uint) (models.User, error) {
	return uc.userRepo.FindOne(id)
}

func (uc *UserUseCase) Register(user models.User) (models.User, error) {
	existUser, err := uc.userRepo.FindByEmail(user.Email)
	if err != nil {
		return models.User{}, err
	}
	if existUser.ID != 0 {
		return models.User{}, errors.New("User Already Exist")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return models.User{}, err
	}

	user.Password = string(encryptedPassword)

	return uc.userRepo.CreateOne(user)
}
