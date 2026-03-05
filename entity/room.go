package entity

import "time"

type Room struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	RoomNumber string    `gorm:"type:varchar(20);not null" json:"room_number"`
	RoomType   string    `gorm:"type:varchar(50)" json:"room_type"`
	Price      float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Capacity   int       `gorm:"not null" json:"capacity"`
	Status     string    `gorm:"type:varchar(20);default:available" json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
