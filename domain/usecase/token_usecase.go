package domain

import (
	"time"

	"datawow/book-list/config"
	"datawow/book-list/models"

	"github.com/golang-jwt/jwt/v5"
)

type ITokenUseCase interface {
	CreateAccessToken(user models.User) (accessToken string, expired int64, err error)
	CreateRefreshToken(user models.User) (refreshToken string, err error)
}

type TokenUseCase struct {
	config *config.Config
}

func NewTokenUseCase(config *config.Config) *TokenUseCase {
	return &TokenUseCase{
		config: config,
	}
}

func (uc *TokenUseCase) CreateAccessToken(user models.User) (t string, expired int64, err error) {
	exp := time.Now().Add(time.Hour * models.ExpireCount)
	claims := &models.JwtCustomClaims{
		Name: user.Name,
		ID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	expired = exp.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err = token.SignedString([]byte(uc.config.Auth.AccessSecret))
	if err != nil {
		return
	}

	return
}

func (uc *TokenUseCase) CreateRefreshToken(user models.User) (t string, err error) {
	claimsRefresh := &models.JwtCustomRefreshClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * models.ExpireRefreshCount)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	rt, err := refreshToken.SignedString([]byte(uc.config.Auth.RefreshSecret))
	if err != nil {
		return "", err
	}
	return rt, err
}
