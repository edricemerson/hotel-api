package repository

import (
	"context"
	"hotel-api/entity"

	"gorm.io/gorm"
)

type BookingRepository struct {
	*gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{db.Table("project_bookings")}
}

func (r *BookingRepository) CreateBooking(booking *entity.Booking) error {

	return r.DB.
		WithContext(context.Background()).
		Create(booking).
		Error
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
