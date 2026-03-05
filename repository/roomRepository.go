package repository

import (
	"context"

	"hotel-api/entity"

	"gorm.io/gorm"
)

type RoomRepository struct {
	*gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db.Table("project_rooms")}
}

func (r *RoomRepository) FindAll() ([]entity.Room, error) {

	var rooms []entity.Room

	err := r.DB.WithContext(context.Background()).
		Find(&rooms).
		Error

	return rooms, err
}

func (r *RoomRepository) FindByID(id string) (entity.Room, error) {

	var room entity.Room

	err := r.DB.WithContext(context.Background()).
		Where("id = ?", id).
		First(&room).
		Error

	return room, err
}

func (r *RoomRepository) Create(room *entity.Room) error {

	return r.DB.WithContext(context.Background()).
		Create(room).
		Error
}

func (r *RoomRepository) Update(id string, room *entity.Room) error {

	return r.DB.WithContext(context.Background()).
		Where("id = ?", id).
		Updates(room).
		Error
}

func (r *RoomRepository) Delete(id string) error {

	return r.DB.WithContext(context.Background()).
		Where("id = ?", id).
		Delete(&entity.Room{}).
		Error
}
