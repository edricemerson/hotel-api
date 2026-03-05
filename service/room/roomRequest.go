package room

import (
	"hotel-api/entity"
)

type Repository interface {
	FindAll() ([]entity.Room, error)
	FindByID(id string) (entity.Room, error)
	Create(room *entity.Room) error
	Update(id string, room *entity.Room) error
	Delete(id string) error
}
