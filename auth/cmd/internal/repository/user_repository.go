package repository

import (
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/model"
	"gorm.io/gorm"
)


type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *model.User) error {
	user.BeforeCreate()
	return r.db.Create(user).Error
}