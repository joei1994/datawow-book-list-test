package handler

import (
	"datawow/book-list/config"
	domain "datawow/book-list/domain/usecase"
	"datawow/book-list/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	BookHandler *BookHandler
	UserHandler *UserHandler
}

func NewHandler(usecase domain.UseCase) *Handler {
	return &Handler{
		BookHandler: NewBookHandler(usecase.BookUseCase),
		UserHandler: NewUserHandler(usecase.UserUseCase),
	}
}

func (h *Handler) Routes(authConfig config.AuthConfig, server *echo.Echo) {
	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey: []byte(authConfig.AccessSecret),
	}

	apiGroup := server.Group("/api")
	apiGroup.Use(echojwt.WithConfig(jwtConfig))
	apiGroup.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(100)))

	h.BookHandler.Routes(apiGroup)

	authGroup := server.Group("/auth")
	h.UserHandler.Routes(authGroup)
}
