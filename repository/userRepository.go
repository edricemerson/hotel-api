package repository

import (
	"context"

	"hotel-api/entity"

	"gorm.io/gorm"
)

type GormRepository struct {
	*gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db.Table("project_users"),
	}
}

func (r *GormRepository) Create(u *entity.User) (err error) {
	ctx := context.Background()

	return r.DB.WithContext(ctx).
		Create(u).
		Error
}

func (r *GormRepository) FindByEmail(email string) (u entity.User, err error) {
	ctx := context.Background()

	err = r.DB.WithContext(ctx).
		Where("email = ?", email).
		First(&u).
		Error

	return
}

func (r *GormRepository) FindByPhone(phone string) (u entity.User, err error) {
	ctx := context.Background()

	err = r.DB.WithContext(ctx).
		Where("phone = ?", phone).
		First(&u).
		Error

	return
}

func (r *GormRepository) FindByName(name string) (u entity.User, err error) {
	ctx := context.Background()

	err = r.DB.WithContext(ctx).
		Where("name = ?", name).
		First(&u).
		Error

	return
}

func (r *GormRepository) FindByID(id string) (u entity.User, err error) {
	ctx := context.Background()

	err = r.DB.WithContext(ctx).
		Where("id = ?", id).
		First(&u).
		Error

	return
}
