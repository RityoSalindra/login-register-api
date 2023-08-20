package main

import (
	"net/http"

	"login-register/config"
	"login-register/controller"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
)

var (
	db *gorm.DB
)

func restricted(c echo.Context) error {
	email := c.Get("email").(string)
	return c.String(http.StatusOK, "Welcome "+email+"!")
}

func main() {
	e := echo.New()

	// Connect To Database
	config.InitDB()

	authController := controller.NewAuthController(db)

	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)

	r := e.Group("/restricted")
	r.Use(controller.JWTMiddleware())
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":8080"))
}
