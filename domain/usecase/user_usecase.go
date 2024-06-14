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
	Login(email, password string) (accessToken, refreshToken string, exp int64, err error)
}

type UserUseCase struct {
	userRepo repository.IUserRepository
	tokenUc  ITokenUseCase
}

func NewUserUseCase(userRepo repository.IUserRepository, tokenUc ITokenUseCase) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
		tokenUc:  tokenUc,
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

func (uc *UserUseCase) Login(email, password string) (accessToken, refreshToken string, expired int64, err error) {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return
	}

	if user.ID == 0 || (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil) {
		return "", "", 0, errors.New(models.INVALID_CREDENTIALS)
	}

	accessToken, expired, err = uc.tokenUc.CreateAccessToken(user)
	if err != nil {
		return
	}

	refreshToken, err = uc.tokenUc.CreateRefreshToken(user)
	if err != nil {
		return
	}

	return accessToken, refreshToken, expired, err
}
