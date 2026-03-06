package booking

import (
	"errors"
	"testing"
	"time"

	"hotel-api/entity"
	"hotel-api/service/room"
	"hotel-api/service/user"
	"hotel-api/util"

	"github.com/golang/mock/gomock"
)

func TestCreateBookingSuccess(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockRoomRepo := room.NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)

	service := NewService(mockRepo, mockRoomRepo, mockUserRepo)

	// mock email
	original := util.SendEmail
	defer func() { util.SendEmail = original }()

	util.SendEmail = func(to, subject, body string) error {
		return nil
	}

	checkIn := time.Now()
	checkOut := checkIn.AddDate(0, 0, 2)

	booking := &entity.Booking{
		UserID:   1,
		RoomID:   1,
		CheckIn:  checkIn,
		CheckOut: checkOut,
	}

	roomData := entity.Room{
		ID:         1,
		RoomNumber: "101",
		RoomType:   "Deluxe",
		Price:      100,
	}

	userData := entity.User{
		ID:    1,
		Email: "test@mail.com",
	}

	mockRoomRepo.EXPECT().
		FindByID("1").
		Return(roomData, nil)

	mockRepo.EXPECT().
		CheckRoomAvailability(1, gomock.Any(), gomock.Any()).
		Return(true, nil)

	mockRepo.EXPECT().
		CreateBooking(gomock.Any()).
		Return(nil)

	mockUserRepo.EXPECT().
		FindByID("1").
		Return(userData, nil)

	mockRoomRepo.EXPECT().
		Update("1", gomock.Any()).
		Return(nil)

	err := service.CreateBooking(booking)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if booking.TotalPrice != 200 {
		t.Errorf("expected total price 200, got %f", booking.TotalPrice)
	}
}

func TestCreateBookingInvalidDate(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockRoomRepo := room.NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)

	service := NewService(mockRepo, mockRoomRepo, mockUserRepo)

	checkIn := time.Now()
	checkOut := checkIn.AddDate(0, 0, -1)

	booking := &entity.Booking{
		CheckIn:  checkIn,
		CheckOut: checkOut,
	}

	err := service.CreateBooking(booking)

	if err == nil {
		t.Fatalf("expected error for invalid date")
	}
}

func TestCreateBookingRoomNotFound(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockRoomRepo := room.NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)

	service := NewService(mockRepo, mockRoomRepo, mockUserRepo)

	checkIn := time.Now()
	checkOut := checkIn.AddDate(0, 0, 2)

	booking := &entity.Booking{
		RoomID:   1,
		CheckIn:  checkIn,
		CheckOut: checkOut,
	}

	mockRoomRepo.EXPECT().
		FindByID("1").
		Return(entity.Room{}, errors.New("not found"))

	err := service.CreateBooking(booking)

	if err == nil {
		t.Fatalf("expected room not found error")
	}
}

func TestCreateBookingRoomNotAvailable(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockRoomRepo := room.NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)

	service := NewService(mockRepo, mockRoomRepo, mockUserRepo)

	checkIn := time.Now()
	checkOut := checkIn.AddDate(0, 0, 2)

	booking := &entity.Booking{
		RoomID:   1,
		CheckIn:  checkIn,
		CheckOut: checkOut,
	}

	roomData := entity.Room{
		ID:    1,
		Price: 100,
	}

	mockRoomRepo.EXPECT().
		FindByID("1").
		Return(roomData, nil)

	mockRepo.EXPECT().
		CheckRoomAvailability(1, gomock.Any(), gomock.Any()).
		Return(false, nil)

	err := service.CreateBooking(booking)

	if err == nil {
		t.Fatalf("expected room unavailable error")
	}
}

func TestGetMyBookings(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockRoomRepo := room.NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)

	service := NewService(mockRepo, mockRoomRepo, mockUserRepo)

	bookings := []entity.Booking{
		{ID: 1, UserID: 1},
	}

	mockRepo.EXPECT().
		GetMyBookings(1).
		Return(bookings, nil)

	result, err := service.GetMyBookings(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("expected 1 booking")
	}
}

func TestDeleteBooking(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockRoomRepo := room.NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)

	service := NewService(mockRepo, mockRoomRepo, mockUserRepo)

	booking := entity.Booking{
		ID:     1,
		RoomID: 1,
	}

	mockRepo.EXPECT().
		GetBookingByID("1").
		Return(booking, nil)

	mockRepo.EXPECT().
		DeleteBooking("1").
		Return(nil)

	mockRoomRepo.EXPECT().
		Update("1", gomock.Any()).
		Return(nil)

	err := service.DeleteBooking("1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
