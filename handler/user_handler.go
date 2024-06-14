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

func (h *UserHandler) Routes(r *echo.Group) {
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}

// Register godoc
//
//	@Summary		Register a user
//	@Description	Register a user
//	@ID				user-register
//	@Tags			Auth Actions
//	@Accept			json
//	@Produce		json
//	@Param			params	body		requests.RegisterRequest	true	"User's credentials"
//	@Success		200		{object}	models.User
//	@Failure		401		{object}	responses.Error
//	@Router			/auth/register [post]
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

// Login godoc
//
//	@Summary		Authenticate a user
//	@Description	Perform user login
//	@ID				user-login
//	@Tags			Auth Actions
//	@Accept			json
//	@Produce		json
//	@Param			params	body		requests.LoginRequest	true	"User's credentials"
//	@Success		200		{object}	responses.LoginResponse
//	@Failure		401		{object}	responses.Error
//	@Router			/auth/login [post]
func (h *UserHandler) Login(c echo.Context) error {
	payload := requests.LoginRequest{}

	if err := c.Bind(&payload); err != nil {
		return err
	}

	if err := payload.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	accessToken, refreshToken, expired, err := h.userUseCase.Login(payload.Email, payload.Password)
	if err != nil && err.Error() == models.INVALID_CREDENTIALS {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid Credentials")
	}

	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Loin Failed")
	}

	resp := responses.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Exp:          expired,
	}

	return responses.Response(c, http.StatusOK, resp)
}

func (h *UserHandler) toUser(payload requests.RegisterRequest) models.User {
	return models.User{
		Email:    payload.Email,
		Password: payload.Password,
		Name:     payload.Name,
	}
}
