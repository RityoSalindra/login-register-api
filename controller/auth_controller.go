package controller

import (
	"net/http"
	"time"

	"login-register/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	DB *gorm.DB
}

func generateToken(email string) (string, error) {
	claims := &model.JWTClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret_key"))
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ac *AuthController) Register(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := model.User{
		Email:    email,
		Password: password,
	}
	ac.DB.Create(&user)

	return c.String(http.StatusCreated, "User registered successfully")
}

func (ac *AuthController) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	var user model.User
	ac.DB.Where("email = ?", email).First(&user)

	if user.ID == 0 || user.Password != password {
		return echo.ErrUnauthorized
	}

	token, err := generateToken(email)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
