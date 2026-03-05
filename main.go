package main

import (
	"log"
	"os"

	"hotel-api/handler"
	"hotel-api/repository"
	"hotel-api/service/room"
	"hotel-api/service/user"
	"hotel-api/util"

	"github.com/labstack/echo/v4"
)

func main() {
	// connect database
	db := util.ConnectDB()

	userRepo := repository.NewGormRepository(db)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	roomRepo := repository.NewRoomRepository(db)
	roomService := room.NewService(roomRepo)
	roomHandler := handler.NewRoomHandler(roomService)

	e := echo.New()

	//public route
	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)

	// admin route
	admin := e.Group("")
	admin.Use(util.JWTMiddleware)
	admin.Use(util.AdminOnly)

	admin.POST("/rooms", roomHandler.CreateRoom)
	admin.PUT("/rooms/:id", roomHandler.UpdateRoom)
	admin.DELETE("/rooms/:id", roomHandler.DeleteRoom)

	// authenticated route
	auth := e.Group("")
	auth.Use(util.JWTMiddleware)

	auth.GET("/rooms", roomHandler.GetRooms)
	auth.GET("/rooms/:id", roomHandler.GetRoomByID)
	auth.POST("", bookingHandler.CreateBooking)
	auth.GET("", bookingHandler.GetMyBookings)

	port := os.Getenv("PORT")

	log.Println("Server running on port", port)

	e.Logger.Fatal(e.Start(":" + port))
}
