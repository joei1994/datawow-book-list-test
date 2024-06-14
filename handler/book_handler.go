package handler

import (
	"net/http"
	"strconv"

	domain "datawow/book-list/domain/usecase"
	"datawow/book-list/models"
	"datawow/book-list/requests"
	"datawow/book-list/responses"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookUseCase domain.IBookUseCase
}

func NewBookHandler(bookUseCase domain.IBookUseCase) *BookHandler {
	return &BookHandler{
		bookUseCase: bookUseCase,
	}
}

func (h *BookHandler) Routes(r *echo.Group) {
	r.POST("/books", h.CreateBook)
	r.GET("/books/:id", h.GetBook)
	r.GET("/books", h.GetBooks)
	r.DELETE("/books/:id", h.DeleteBook)
	r.PUT("/books/:id", h.UpdateBook)
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	payload := requests.BookPayload{}
	if err := c.Bind(&payload); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := payload.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	book, err := h.bookUseCase.CreateBook(h.toBook(payload, claims.Name))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.Response(c, http.StatusOK, book)
}

func (h *BookHandler) GetBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	book, err := h.bookUseCase.GetBook(uint(id))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.Response(c, http.StatusOK, book)
}

func (h *BookHandler) GetBooks(c echo.Context) error {
	ids := []uint{}
	for _, id := range c.QueryParams()["id"] {
		if idInt, err := strconv.Atoi(id); err == nil {
			ids = append(ids, uint(idInt))
		}
	}

	books, err := h.bookUseCase.GetBooks(ids)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.Response(c, http.StatusOK, books)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.bookUseCase.DeleteBook(uint(id)); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.MessageResponse(c, http.StatusOK, models.DELETE_SUCCESS)
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	reqPayload := requests.BookPayload{}
	if err := c.Bind(&reqPayload); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := reqPayload.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	book, err := h.bookUseCase.UpdateBook(uint(id), h.toBook(reqPayload, claims.Name))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.Response(c, http.StatusOK, book)
}

func (h *BookHandler) toBook(req requests.BookPayload, createdBy string) models.Book {
	return models.Book{
		Title:     req.Title,
		Author:    req.Author,
		Genre:     req.Genre,
		Year:      req.Year,
		Tag:       req.Tag,
		CreatedBy: createdBy,
	}
}
