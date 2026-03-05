package main

import (
	"log"
	"os"

	"hotel-api/handler"
	"hotel-api/repository"
	"hotel-api/service/user"
	"hotel-api/util"

	"github.com/labstack/echo/v4"
)

func main() {
	// connect database
	db := util.ConnectDB()

	repo := repository.NewGormRepository(db)
	service := user.NewService(repo)
	handler := handler.NewUserHandler(service)

	e := echo.New()

	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)

	port := os.Getenv("PORT")

	log.Println("Server running on port", port)

	e.Logger.Fatal(e.Start(":" + port))
}
