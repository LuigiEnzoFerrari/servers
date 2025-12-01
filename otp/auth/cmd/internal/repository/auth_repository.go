package repository

import "gorm.io/gorm"

type AuthRepository struct {
	db *gorm.DB
}

