package repository

import (
	"context"
	"errors"
	"hotel-api/entity"
	"time"

	"gorm.io/gorm"
)

type BookingRepository struct {
	*gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{db.Table("project_bookings")}
}

func (r *BookingRepository) CreateBooking(booking *entity.Booking) error {

	return r.DB.Transaction(func(tx *gorm.DB) error {

		var count int64

		err := tx.Model(&entity.Booking{}).
			Where("room_id = ?", booking.RoomID).
			Where("booking_status = ?", "confirmed").
			Where("check_in < ? AND check_out > ?", booking.CheckOut, booking.CheckIn).
			Count(&count).Error

		if err != nil {
			return err
		}

		if count > 0 {
			return errors.New("room already booked for these dates")
		}

		return tx.Create(booking).Error
	})
}

func (r *BookingRepository) GetMyBookings(userID int) ([]entity.Booking, error) {

	var bookings []entity.Booking

	err := r.DB.
		WithContext(context.Background()).
		Preload("Room").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&bookings).Error

	return bookings, err
}

func (r *BookingRepository) GetBookingByID(id string) (entity.Booking, error) {

	var booking entity.Booking

	err := r.DB.
		WithContext(context.Background()).
		Preload("Room").
		Where("id = ?", id).
		First(&booking).
		Error

	return booking, err
}

func (r *BookingRepository) UpdateBooking(id string, booking *entity.Booking) error {

	return r.DB.
		WithContext(context.Background()).
		Model(&entity.Booking{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"check_in":  booking.CheckIn,
			"check_out": booking.CheckOut,
		}).Error
}

func (r *BookingRepository) DeleteBooking(id string) error {

	return r.DB.
		WithContext(context.Background()).
		Delete(&entity.Booking{}, "id = ?", id).
		Error
}

func (r *BookingRepository) CheckRoomAvailability(roomID int, checkIn, checkOut time.Time) (bool, error) {

	var count int64

	err := r.DB.
		WithContext(context.Background()).
		Model(&entity.Booking{}).
		Where("room_id = ?", roomID).
		Where("booking_status = ?", "confirmed").
		Where("check_in < ? AND check_out > ?", checkOut, checkIn).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count == 0, nil
}
