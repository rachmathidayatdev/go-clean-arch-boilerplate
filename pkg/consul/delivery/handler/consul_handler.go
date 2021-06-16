package handler

import (
	"net/http"

	"github.com/go-clean-arch-boilerplate/library"
	"github.com/labstack/echo/v4"
)

// HealthCheck func
// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Accept */*
// @Produce json
// @Success 200 {object} library.HTTPResponse
// @Router /health [get]
func (handler *initHandler) HealthCheck(c echo.Context) error {
	response := library.HTTPResponse{
		StatusCode: "200",
		Message:    "success",
		Data:       nil,
	}

	return c.JSON(http.StatusOK, response)
}
