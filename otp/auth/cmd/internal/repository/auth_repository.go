package repository

import (
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/domain"
	"gorm.io/gorm"
)

type PostgresAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *PostgresAuthRepository {
	return &PostgresAuthRepository{db: db}
}

func (r *PostgresAuthRepository) Save(user *domain.Auth) error {
	user.BeforeCreate()
	return r.db.Create(user).Error
}
