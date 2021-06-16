package handler

import (
	"github.com/go-clean-arch-boilerplate/pkg/consul"
	consulRepository "github.com/go-clean-arch-boilerplate/pkg/consul/repository"
	consulUsecase "github.com/go-clean-arch-boilerplate/pkg/consul/usecase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

//initHandler struct
type initHandler struct {
	Usecase *useCase
}

type useCase struct {
	Consul consul.Usecase
}

// Init func
func Init(e *echo.Echo, db *gorm.DB) {
	consulRepository := consulRepository.Repository(db)
	consulUsecase := consulUsecase.Usecase(consulRepository)

	useCase := &useCase{
		Consul: consulUsecase,
	}

	handler := &initHandler{
		Usecase: useCase,
	}

	e.GET("/health", handler.HealthCheck)
}
