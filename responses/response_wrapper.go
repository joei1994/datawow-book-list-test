package responses

import (
	"github.com/labstack/echo/v4"
)

type Error struct {
	Code  int    `json:"code" example:"404"`
	Error string `json:"error" example:"Data Not Found"`
}

type Data struct {
	Code    int    `json:"code" example:"202"`
	Message string `json:"message" example:"Create Success"`
}

func Response(c echo.Context, statusCode int, data interface{}) error {
	return c.JSON(statusCode, data)
}

func MessageResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Data{
		Code:    statusCode,
		Message: message,
	})
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}
