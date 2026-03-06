package util

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "authorization header missing",
			})
		}

		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid authorization format",
			})
		}

		tokenString := authHeader[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "unexpected signing method")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token claims",
			})
		}

		c.Set("user", claims)

		return next(c)
	}
}

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		user := c.Get("user")
		if user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "user not found in token",
			})
		}

		claims, ok := user.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token claims",
			})
		}

		role, ok := claims["role"].(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "role missing in token",
			})
		}

		if role != "admin" {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "admin access required",
			})
		}

		return next(c)
	}
}
