package domain

import (
	"time"
	"github.com/google/uuid"
)

type Auth struct {
	Id           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email        string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type AuthRepository interface {
	Save(auth *Auth) error
	FindByEmail(email string) (*Auth, error)
}

type AuthRequestDTO struct {
	Email    string
	Password string
}

func (u *Auth) BeforeCreate() (err error) {
	u.Id = uuid.New()
	return
}

type PasswordForgotEvent struct {
	Email string 
}

type AuthPublish interface {
	PublishPasswordForgotEvent(event PasswordForgotEvent) error
}
