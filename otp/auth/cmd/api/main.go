package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/repository"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	repository.NewAuthRepository(db)



	fmt.Println("Hello, World!")
}
