package booking

import (
	"hotel-api/entity"
	"time"
)

type Repository interface {
	CreateBooking(booking *entity.Booking) error
	GetMyBookings(userID int) ([]entity.Booking, error)
	GetBookingByID(id string) (entity.Booking, error)
	UpdateBooking(id string, booking *entity.Booking) error
	DeleteBooking(id string) error
	CheckRoomAvailability(roomID int, checkIn, checkOut time.Time) (bool, error)
}
