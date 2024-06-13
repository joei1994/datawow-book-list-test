package handler

import (
	"net/http"

	domain "datawow/book-list/domain/usecase"

	"datawow/book-list/models"
	"datawow/book-list/requests"
	"datawow/book-list/responses"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase domain.IUserUseCase
}

func NewUserHandler(userUseCase domain.IUserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) RegisterRoutes(r *echo.Group) {
	r.POST("/register", h.Register)
}

func (h *UserHandler) Register(c echo.Context) error {
	payload := requests.RegisterRequest{}

	if err := c.Bind(&payload); err != nil {
		return err
	}

	if err := payload.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	user, err := h.userUseCase.Register(h.toUser(payload))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.Response(c, http.StatusOK, user)
}

func (h *UserHandler) toUser(payload requests.RegisterRequest) models.User {
	return models.User{
		Email:    payload.Email,
		Password: payload.Password,
		Name:     payload.Name,
	}
}
