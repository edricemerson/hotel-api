package user

import (
	"hotel-api/entity"
)

type Repository interface {
	Create(user *entity.User) (err error)
	FindByEmail(email string) (user entity.User, err error)
	FindByPhone(phone string) (entity.User, error)
	FindByName(name string) (entity.User, error)
}
