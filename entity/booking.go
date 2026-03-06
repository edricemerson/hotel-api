package entity

import "time"

type Booking struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        int       `gorm:"not null" json:"user_id"`
	RoomID        int       `gorm:"not null" json:"room_id"`
	CheckIn       time.Time `gorm:"type:date;not null" json:"check_in"`
	CheckOut      time.Time `gorm:"type:date;not null" json:"check_out"`
	TotalPrice    float64   `gorm:"type:decimal(10,2);not null" json:"-"`
	BookingStatus string    `gorm:"type:varchar(20);default:confirmed" json:"-"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Room Room `gorm:"foreignKey:RoomID" json:"room"`
}

func (Booking) TableName() string {
	return "project_bookings"
}

type BookingRequest struct {
	RoomID   int    `json:"room_id"`
	CheckIn  string `json:"check_in"`
	CheckOut string `json:"check_out"`
}
