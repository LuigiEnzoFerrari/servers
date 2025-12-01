package domain

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Username string
	Password string
}
