package handler

import (
	domain "datawow/book-list/domain/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	BookHandler *BookHandler
}

func NewHandler(usecase domain.UseCase) *Handler {
	return &Handler{
		BookHandler: NewBookHandler(usecase.BookUseCase),
	}
}

func (h *Handler) RegisterRoutes(server *echo.Echo) {
	group := server.Group("/api")

	h.BookHandler.RegisterRoutes(group)
}
