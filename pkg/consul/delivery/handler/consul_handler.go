package handler

import (
	"net/http"

	"github.com/go-clean-arch-boilerplate/library"
	"github.com/labstack/echo"
)

//HealthCheck func
func (handler *initHandler) HealthCheck(c echo.Context) error {
	response := library.HTTPResponse{
		StatusCode: "200",
		Message:    "success",
		Data:       nil,
	}

	return c.JSON(http.StatusOK, response)
}
