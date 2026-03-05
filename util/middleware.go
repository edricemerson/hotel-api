package util

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		tokenString := c.Request().Header.Get("Authorization")

		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing token",
			})
		}

		tokenString = tokenString[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token",
			})
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set("user", claims)

		return next(c)
	}
}

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		user := c.Get("user").(map[string]interface{})

		role := user["role"].(string)

		if role != "admin" {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "admin access required",
			})
		}

		return next(c)
	}
}
