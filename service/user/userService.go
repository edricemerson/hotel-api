package user

import (
	"errors"
	"hotel-api/entity"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo Repository
}

type Service interface {
	Register(name string, email string, password string) error
	Login(email string, password string) (entity.User, error)
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Register(name string, email string, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := entity.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.repo.Create(user)
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
