package repository

import (
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"` 
	Username     string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	
	CreatedAt    time.Time 
	UpdatedAt    time.Time 
}

func (a *Auth) beforeCreate() (err error) {
    a.ID = uuid.New()
    return
}

