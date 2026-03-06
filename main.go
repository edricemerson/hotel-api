// @title Hotel API
// @version 1.0
// @description A comprehensive Hotel Management API with user authentication, room management, and booking system
// @host
// @basePath /
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

	// load env (only for local)
	if err := godotenv.Load(); err != nil {
		log.Println("Running without .env (Railway environment)")
	}

	// connect database
	db := util.ConnectDB()

	// repository
	userRepo := repository.NewGormRepository(db)
	roomRepo := repository.NewRoomRepository(db)
	bookRepo := repository.NewBookingRepository(db)

	// service
	userService := user.NewService(userRepo)
	roomService := room.NewService(roomRepo)
	bookingService := booking.NewService(bookRepo, roomRepo, userRepo)

	// handler
	userHandler := handler.NewUserHandler(userService)
	roomHandler := handler.NewRoomHandler(roomService)
	bookingHandler := handler.NewBookingHandler(bookingService)

	// echo instance
	e := echo.New()

	// health check
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "Hotel API running",
		})
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)

	auth := e.Group("")
	auth.Use(util.JWTMiddleware)

	auth.GET("/rooms", roomHandler.GetRooms)
	auth.GET("/rooms/:id", roomHandler.GetRoomByID)

	auth.POST("/bookings", bookingHandler.CreateBooking)
	auth.GET("/bookings", bookingHandler.GetMyBookings)

	admin := e.Group("")
	admin.Use(util.JWTMiddleware)
	admin.Use(util.AdminOnly)

	admin.POST("/rooms", roomHandler.CreateRoom)
	admin.PUT("/rooms/:id", roomHandler.UpdateRoom)
	admin.DELETE("/rooms/:id", roomHandler.DeleteRoom)

	admin.GET("/bookings/:id", bookingHandler.GetBookingByID)
	admin.PUT("/bookings/:id", bookingHandler.UpdateBooking)
	admin.DELETE("/bookings/:id", bookingHandler.DeleteBooking)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)

	e.Logger.Fatal(e.Start(":" + port))
}
