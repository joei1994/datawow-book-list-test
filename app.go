package application

import (
	"log"

	"datawow/book-list/config"
	domain "datawow/book-list/domain/usecase"
	"datawow/book-list/handler"
	"datawow/book-list/infrastructure"
	"datawow/book-list/repository"

	"github.com/labstack/echo/v4"
)

func Start(config *config.Config) {
	server := echo.New()

	infras := infrastructure.NewInfrastructure(config)
	repos := repository.NewRepository(infras)
	ucs := domain.NewUseCase(config, repos)
	handlers := handler.NewHandler(ucs)

	handlers.RegisterRoutes(server)

	err := server.Start(":" + config.Http.Port)
	if err != nil {
		log.Fatal("Port already used: ", err)
	}
}
