package controller

import (
	"login-register/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return echo.ErrUnauthorized
			}

			token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret_key"), nil
			})
			if err != nil {
				return echo.ErrUnauthorized
			}

			claims, ok := token.Claims.(*model.JWTClaims)
			if !ok || !token.Valid {
				return echo.ErrUnauthorized
			}

			c.Set("email", claims.Email)
			return next(c)
		}
	}
}
