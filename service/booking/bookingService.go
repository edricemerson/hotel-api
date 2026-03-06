package booking

import (
	"errors"
	"hotel-api/entity"
	"hotel-api/service/room"
	"strconv"
)

type service struct {
	repo     Repository
	roomRepo room.Repository
}

type Service interface {
	CreateBooking(booking *entity.Booking) error
	GetMyBookings(userID int) ([]entity.Booking, error)
	GetBookingByID(id string) (entity.Booking, error)
	UpdateBooking(id string, booking *entity.Booking) error
	DeleteBooking(id string) error
}

func NewService(r Repository, roomRepo room.Repository) Service {
	return &service{
		repo:     r,
		roomRepo: roomRepo,
	}
}

func (s *service) CreateBooking(booking *entity.Booking) error {

	if !booking.CheckOut.After(booking.CheckIn) {
		return errors.New("check_out must be after check_in")
	}

	available, err := s.repo.CheckRoomAvailability(
		booking.RoomID,
		booking.CheckIn,
		booking.CheckOut,
	)

	if err != nil {
		return err
	}

	if !available {
		return errors.New("room already booked for selected dates")
	}

	booking.BookingStatus = "confirmed"

	err = s.repo.CreateBooking(booking)
	if err != nil {
		return err
	}

	room := entity.Room{
		Status: "unavailable",
	}

	err = s.roomRepo.Update(
		strconv.Itoa(booking.RoomID),
		&room,
	)

	return err
}

func (s *service) GetMyBookings(userID int) ([]entity.Booking, error) {
	return s.repo.GetMyBookings(userID)
}

func (s *service) GetBookingByID(id string) (entity.Booking, error) {
	return s.repo.GetBookingByID(id)
}

func (s *service) UpdateBooking(id string, booking *entity.Booking) error {

	if booking.CheckOut.Before(booking.CheckIn) {
		return errors.New("check_out must be after check_in")
	}

	return s.repo.UpdateBooking(id, booking)
}

func (s *service) DeleteBooking(id string) error {
	return s.repo.DeleteBooking(id)
}
