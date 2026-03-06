package room

import (
	"errors"
	"testing"

	"hotel-api/entity"

	"github.com/golang/mock/gomock"
)

func TestCreateRoomSuccess(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	service := NewService(mockRepo)

	room := &entity.Room{
		RoomNumber: "101",
		RoomType:   "Deluxe",
		Status:     "available",
	}

	mockRepo.EXPECT().
		FindByRoomNumber("101").
		Return(entity.Room{}, errors.New("not found"))

	mockRepo.EXPECT().
		Create(gomock.Any()).
		Return(nil)

	err := service.Create(room)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if room.Capacity != 2 {
		t.Errorf("expected capacity 2, got %d", room.Capacity)
	}
}

func TestCreateRoomAlreadyExists(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	service := NewService(mockRepo)

	room := &entity.Room{
		RoomNumber: "101",
		RoomType:   "Deluxe",
		Status:     "available",
	}

	mockRepo.EXPECT().
		FindByRoomNumber("101").
		Return(entity.Room{ID: 1}, nil)

	err := service.Create(room)

	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

func TestCreateInvalidRoomType(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	service := NewService(mockRepo)

	room := &entity.Room{
		RoomNumber: "101",
		RoomType:   "Standard",
		Status:     "available",
	}

	mockRepo.EXPECT().
		FindByRoomNumber("101").
		Return(entity.Room{}, errors.New("not found"))

	err := service.Create(room)

	if err == nil {
		t.Fatalf("expected error for invalid room type")
	}
}

func TestCreateInvalidStatus(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	service := NewService(mockRepo)

	room := &entity.Room{
		RoomNumber: "101",
		RoomType:   "Deluxe",
		Status:     "maintenance",
	}

	mockRepo.EXPECT().
		FindByRoomNumber("101").
		Return(entity.Room{}, errors.New("not found"))

	err := service.Create(room)

	if err == nil {
		t.Fatalf("expected status error")
	}
}

func TestGetRooms(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	service := NewService(mockRepo)

	rooms := []entity.Room{
		{ID: 1, RoomNumber: "101"},
	}

	mockRepo.EXPECT().
		FindAll().
		Return(rooms, nil)

	result, err := service.GetRooms()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("expected 1 room")
	}
}
