package repository

import (
	"errors"
	"fmt"

	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/domain"
	"github.com/jackc/pgx/v5/pgconn"
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
	if err := r.db.Create(dbAuth).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.ErrConflict
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.ErrConflict
		}
		return fmt.Errorf("db: save user: %w", err)
	}

	return nil
}

func (r *PostgresAuthRepository) FindByUsername(username string) (*domain.Auth, error) {
	var auth Auth
	if err := r.db.Where("username = ?", username).First(&auth).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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
