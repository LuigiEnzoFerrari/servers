package repository

import (
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/domain"
	"gorm.io/gorm"
)


type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(user *domain.User) error {
	user.BeforeCreate()
	return r.db.Create(user).Error
}

func (r *PostgresUserRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}