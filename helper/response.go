package helper

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ResponsePayload struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResOK(c echo.Context, message string, data any) error {
	res := ResponsePayload{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}

func ResErr(c echo.Context, status int, message string) error {
	res := ResponsePayload{
		Status:  status,
		Message: message,
	}
	return c.JSON(status, res)
}

func ResErrHandler(c echo.Context, err error) error {
	if err == gorm.ErrRecordNotFound {
		return ResErr(c, http.StatusNotFound, err.Error())
	}
	return ResErr(c, http.StatusBadRequest, err.Error())
}
