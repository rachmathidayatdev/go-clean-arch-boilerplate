package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-clean-arch-boilerplate/config"
	"github.com/go-clean-arch-boilerplate/router"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//App struct
type App struct {
	Router *echo.Echo
	DB     *gorm.DB
}

func main() {
	error := godotenv.Load()

	if error != nil {
		log.Fatalf("Error loading .env file: %v", error)
	}

	app := App{}
	app.Initialize()
}

//Initialize func
func (app *App) Initialize() {
	connection := config.GetConnection()

	app.Router = echo.New()

	// Middleware
	app.Router.Use(middleware.Logger())
	app.Router.Use(middleware.Recover())

	app.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	router.Route(app.Router, connection)

	// apiServicePort := os.Getenv("API_PORT")
	apiServicePort := os.Getenv("PORT")

	if apiServicePort == "" {
		apiServicePort = os.Getenv("API_PORT")
	}

	http.Handle("/", app.Router)

	log.Printf("API Service listening on port %v", apiServicePort)

	apiServer := &http.Server{
		Addr:         ":" + apiServicePort,
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	err := apiServer.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Println(err.Error())
	}

	defer connection.Close()
}
