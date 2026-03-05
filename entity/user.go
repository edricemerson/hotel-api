package entity

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string    `gorm:"type:text;not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
