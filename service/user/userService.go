package user

import (
	"errors"
	"fmt"
	"regexp"

	"hotel-api/entity"
	"hotel-api/util"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo Repository
}

type Service interface {
	Register(name string, email string, password string, phone string) (entity.User, error)
	Login(email string, password string) (entity.User, error)
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Register(name string, email string, password string, phone string) (entity.User, error) {
	isNumber := regexp.MustCompile(`^[0-9]+$`).MatchString
	if !isNumber(phone) {
		return entity.User{}, errors.New("phone must contain only numbers")
	}

	if len(phone) > 12 {
		return entity.User{}, errors.New("phone must not exceed 12 digits")
	}
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return entity.User{}, errors.New("email already used")
	}
	_, err = s.repo.FindByPhone(phone)
	if err == nil {
		return entity.User{}, errors.New("phone already used")
	}

	_, err = s.repo.FindByName(name)
	if err == nil {
		return entity.User{}, errors.New("name already used")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Phone:    phone,
		Role:     "user",
	}

	err = s.repo.Create(&user)
	if err != nil {
		return entity.User{}, err
	}

	err = util.SendEmail(
		user.Email,
		"Welcome to Hotel API",
		fmt.Sprintf(
			"Hello %s!\n\nYour account has been successfully created.\n\nYou can now login and start booking rooms.\n\nEnjoy your stay!",
			user.Name,
		),
	)

	if err != nil {
		fmt.Println("Email error:", err)
	}

	return user, nil
}

func (s *service) Login(email string, password string) (entity.User, error) {

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return entity.User{}, errors.New("invalid email")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return entity.User{}, errors.New("invalid password")
	}

	return user, nil
}
