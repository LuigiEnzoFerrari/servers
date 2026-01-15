package repository

import (
	"fmt"

	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/domain"
	"gorm.io/gorm"
)


type PostgresAuthRepository struct {
	db *gorm.DB
}

func NewPostgresAuthRepository(db *gorm.DB) *PostgresAuthRepository {
	return &PostgresAuthRepository{db: db}
}

func (r *PostgresAuthRepository) Save(auth *domain.Auth) error {
	dbAuth := Auth{
		Username:     auth.Username,
		PasswordHash: auth.PasswordHash,
		CreatedAt:    auth.CreatedAt,
		UpdatedAt:    auth.UpdatedAt,
	}
	dbAuth.beforeCreate()
	return r.db.Create(dbAuth).Error
}

func (r *PostgresAuthRepository) FindByUsername(username string) (*domain.Auth, error) {
	var auth Auth
	if err := r.db.Where("username = ?", username).First(&auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("db: find user by username: %w", err)
	}
	domainAuth := domain.Auth{
		ID:           auth.ID,
		Username:     auth.Username,
		PasswordHash: auth.PasswordHash,
		CreatedAt:    auth.CreatedAt,
		UpdatedAt:    auth.UpdatedAt,
	}
	return &domainAuth, nil
}
