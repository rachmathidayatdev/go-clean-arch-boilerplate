package repository

import (
	"github.com/go-clean-arch-boilerplate/pkg/consul"
	"github.com/jinzhu/gorm"
)

//dbHandler struct
type dbHandler struct {
	DB *gorm.DB
}

//Repository func
func Repository(DB *gorm.DB) consul.Repository {
	return &dbHandler{
		DB,
	}
}
