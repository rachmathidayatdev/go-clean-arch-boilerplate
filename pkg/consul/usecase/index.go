package usecase

import (
	"github.com/go-clean-arch-boilerplate/pkg/consul"
)

//initData struct
type initData struct {
	Repository *repository
}

type repository struct {
	Consul consul.Repository
}

//Usecase func
func Usecase(consul consul.Repository) consul.Usecase {
	repository := &repository{
		Consul: consul,
	}

	return &initData{
		Repository: repository,
	}
}
