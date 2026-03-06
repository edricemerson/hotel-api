package booking

import (
	"errors"
	"fmt"
	"hotel-api/entity"
	"hotel-api/service/room"
	"hotel-api/service/user"
	"hotel-api/util"
	"strconv"
)

type service struct {
	repo     Repository
	roomRepo room.Repository
	userRepo user.Repository
}

type Service interface {
	CreateBooking(booking *entity.Booking) error
	GetMyBookings(userID int) ([]entity.Booking, error)
	GetBookingByID(id string) (entity.Booking, error)
	UpdateBooking(id string, booking *entity.Booking) error
	DeleteBooking(id string) error
}

func NewService(r Repository, roomRepo room.Repository, userRepo user.Repository) Service {
	return &service{
		repo:     r,
		roomRepo: roomRepo,
		userRepo: userRepo,
	}
}

func (s *service) CreateBooking(booking *entity.Booking) error {

	if !booking.CheckOut.After(booking.CheckIn) {
		return errors.New("check_out must be after check_in")
	}

	roomData, err := s.roomRepo.FindByID(strconv.Itoa(booking.RoomID))
	if err != nil {
		return errors.New("room not found")
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

	nights := int(booking.CheckOut.Sub(booking.CheckIn).Hours() / 24)
	if nights <= 0 {
		return errors.New("invalid booking duration")
	}

	booking.TotalPrice = float64(nights) * roomData.Price
	booking.BookingStatus = "confirmed"

	err = s.repo.CreateBooking(booking)
	if err != nil {
		return err
	}

	user, err := s.userRepo.FindByID(strconv.Itoa(booking.UserID))
	if err != nil {
		return err
	}

	util.SendEmail(
		user.Email,
		"Booking Confirmation",
		fmt.Sprintf(
			"Your booking is confirmed!\n\nRoom Number: %s\nRoom Type: %s\nCheck-in: %s\nCheck-out: %s\nTotal Price: %.2f",
			roomData.RoomNumber,
			roomData.RoomType,
			booking.CheckIn.Format("2006-01-02"),
			booking.CheckOut.Format("2006-01-02"),
			booking.TotalPrice,
		),
	)

	room := entity.Room{
		Status: "unavailable",
	}

	return s.roomRepo.Update(
		strconv.Itoa(booking.RoomID),
		&room,
	)
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
	booking, err := s.repo.GetBookingByID(id)
	if err != nil {
		return err
	}

	err = s.repo.DeleteBooking(id)
	if err != nil {
		return err
	}

	room := entity.Room{
		Status: "available",
	}

	err = s.roomRepo.Update(
		strconv.Itoa(booking.RoomID),
		&room,
	)

	return err
}
