package user

import (
	"errors"
	"testing"

	"hotel-api/entity"

	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestRegister(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	service := NewService(mockRepo)

	name := "John"
	email := "john@mail.com"
	password := "123456"
	phone := "08123456789"

	// email not found
	mockRepo.EXPECT().
		FindByEmail(email).
		Return(entity.User{}, errors.New("not found"))

	// phone not found
	mockRepo.EXPECT().
		FindByPhone(phone).
		Return(entity.User{}, errors.New("not found"))

	// name not found
	mockRepo.EXPECT().
		FindByName(name).
		Return(entity.User{}, errors.New("not found"))

	// create user
	mockRepo.EXPECT().
		Create(gomock.Any()).
		Return(nil)

	_, err := service.Register(name, email, password, phone)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestLoginSuccess(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	service := NewService(mockRepo)

	email := "john@mail.com"
	password := "123456"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := entity.User{
		ID:       1,
		Email:    email,
		Password: string(hashedPassword),
	}

	mockRepo.EXPECT().
		FindByEmail(email).
		Return(user, nil)

	result, err := service.Login(email, password)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Email != email {
		t.Errorf("expected email %s, got %s", email, result.Email)
	}
}

func TestLoginInvalidEmail(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	service := NewService(mockRepo)

	email := "notfound@mail.com"
	password := "123456"

	mockRepo.EXPECT().
		FindByEmail(email).
		Return(entity.User{}, errors.New("not found"))

	_, err := service.Login(email, password)

	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

func TestLoginInvalidPassword(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	service := NewService(mockRepo)

	email := "john@mail.com"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)

	user := entity.User{
		ID:       1,
		Email:    email,
		Password: string(hashedPassword),
	}

	mockRepo.EXPECT().
		FindByEmail(email).
		Return(user, nil)

	_, err := service.Login(email, "wrongpassword")

	if err == nil {
		t.Fatalf("expected password error")
	}
}
