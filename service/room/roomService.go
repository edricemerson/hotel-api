package room

import "hotel-api/entity"

type service struct {
	repo Repository
}

type Service interface {
	GetRooms() ([]entity.Room, error)
	GetRoomByID(id string) (entity.Room, error)
	Create(room entity.Room) error
	Update(id string, room entity.Room) error
	Delete(id string) error
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetRooms() ([]entity.Room, error) {
	return s.repo.FindAll()
}

func (s *service) GetRoomByID(id string) (entity.Room, error) {
	return s.repo.FindByID(id)
}

func (s *service) Create(room entity.Room) error {
	return s.repo.Create(&room)
}

func (s *service) Update(id string, room entity.Room) error {
	return s.repo.Update(id, &room)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
