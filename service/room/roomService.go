package room

import (
	"errors"
	"hotel-api/entity"
)

type service struct {
	repo Repository
}

type Service interface {
	FindByRoomNumber(roomNumber string) (entity.Room, error)
	GetRooms() ([]entity.Room, error)
	GetRoomByID(id string) (entity.Room, error)
	Create(room *entity.Room) error
	Update(id string, room entity.Room) error
	Delete(id string) error
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) FindByRoomNumber(roomNumber string) (entity.Room, error) {
	return s.repo.FindByRoomNumber(roomNumber)
}

func (s *service) GetRooms() ([]entity.Room, error) {
	return s.repo.FindAll()
}

func (s *service) GetRoomByID(id string) (entity.Room, error) {
	return s.repo.FindByID(id)
}

func (s *service) Create(room *entity.Room) error {

	_, err := s.repo.FindByRoomNumber(room.RoomNumber)
	if err == nil {
		return errors.New("room number already exists")
	}

	switch room.RoomType {
	case "Deluxe":
		room.Capacity = 2
	case "Suites":
		room.Capacity = 4
	case "Presidential":
		room.Capacity = 6
	default:
		return errors.New("room_type must be Deluxe, Suites, or Presidential")
	}

	if room.Status != "available" && room.Status != "unavailable" {
		return errors.New("status must be available or unavailable")
	}

	return s.repo.Create(room)
}

func (s *service) Update(id string, room entity.Room) error {

	switch room.RoomType {
	case "Deluxe":
		room.Capacity = 2
	case "Suites":
		room.Capacity = 4
	case "Presidential":
		room.Capacity = 6
	default:
		return errors.New("room_type must be Deluxe, Suites, or Presidential")
	}

	if room.Status != "available" && room.Status != "unavailable" {
		return errors.New("status must be available or unavailable")
	}

	return s.repo.Update(id, &room)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
