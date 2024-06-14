package application

import (
	"log"

	"datawow/book-list/config"
	domain "datawow/book-list/domain/usecase"
	"datawow/book-list/handler"
	"datawow/book-list/infrastructure"
	"datawow/book-list/repository"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

func Start(config *config.Config) {
	server := echo.New()
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	infras := infrastructure.NewInfrastructure(config)
	repos := repository.NewRepository(infras)
	ucs := domain.NewUseCase(config, repos)
	handlers := handler.NewHandler(ucs)

	handlers.Routes(config.Auth, server)

	err := server.Start(":" + config.Http.Port)
	if err != nil {
		log.Fatal("Port already used: ", err)
	}
}
