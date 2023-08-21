package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
