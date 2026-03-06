// @title Hotel API
// @version 1.0
// @description A comprehensive Hotel Management API with user authentication, room management, and booking system
// @host localhost:8080
// @basePath /api
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @schemes http https

package main

import (
	"log"
	"os"

	"hotel-api/handler"
	"hotel-api/repository"
	"hotel-api/service/booking"
	"hotel-api/service/room"
	"hotel-api/service/user"
	"hotel-api/util"

	_ "hotel-api/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// connect database
	db := util.ConnectDB()

	userRepo := repository.NewGormRepository(db)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	roomRepo := repository.NewRoomRepository(db)
	roomService := room.NewService(roomRepo)
	roomHandler := handler.NewRoomHandler(roomService)

	bookRepo := repository.NewBookingRepository(db)
	bookingService := booking.NewService(bookRepo, roomRepo, userRepo)
	bookingHandler := handler.NewBookingHandler(bookingService)

	e := echo.New()

	// Swagger route for Echo v4
	e.GET("/swagger/*", echoSwagger.WrapHandler)

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

	admin.GET("/bookings/:id", bookingHandler.GetBookingByID)
	admin.PUT("/bookings/:id", bookingHandler.UpdateBooking)
	admin.DELETE("/bookings/:id", bookingHandler.DeleteBooking)

	// authenticated route
	auth := e.Group("")
	auth.Use(util.JWTMiddleware)

	auth.GET("/rooms", roomHandler.GetRooms)
	auth.GET("/rooms/:id", roomHandler.GetRoomByID)

	auth.POST("/bookings", bookingHandler.CreateBooking)
	auth.GET("/bookings", bookingHandler.GetMyBookings)

	port := os.Getenv("PORT")

	log.Println("Server running on port", port)

	e.Logger.Fatal(e.Start(":" + port))
}
