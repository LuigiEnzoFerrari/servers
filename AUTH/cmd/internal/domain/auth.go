package domain

import (
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	ID           uuid.UUID
	Username     string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type JwtToken struct {
	Token string
}