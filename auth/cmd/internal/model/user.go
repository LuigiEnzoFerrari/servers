package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"` 
	Username     string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	
	CreatedAt    time.Time 
	UpdatedAt    time.Time 
}

func (u *User) BeforeCreate() (err error) {
    u.ID = uuid.New()
    return
}