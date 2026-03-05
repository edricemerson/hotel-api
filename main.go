package main

import (
	"log"
	"net/http"
	"os"

	"hotel-api/util"

	"github.com/labstack/echo/v4"
)

func main() {
	// connect database
	util.ConnectDB()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Testing API",
		})
	})

	port := os.Getenv("PORT")

	log.Println("Server running on port", port)

	e.Logger.Fatal(e.Start(":" + port))
}
