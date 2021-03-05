package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/go-clean-arch-boilerplate/models"
	consulHandler "github.com/go-clean-arch-boilerplate/pkg/consul/delivery/handler"
)

// Route func
func Route(e *echo.Echo, db *gorm.DB) {

	go func() {
		// dbMigrate(db)

		consulHandler.Init(e, db)
	}()

}

//dbMigrate func
func dbMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(
		&models.Merchant{},
	)

	return db
}
