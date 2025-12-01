package main

import (
	"fmt"

	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/repository"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/service"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)


func main() {

	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=pass dbname=auth port=5433 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	repository := repository.NewAuthRepository(db)
	service := service.NewAuthService(repository)
	authHandler := handler.NewAuthHandler(service)

	gin := gin.Default()
	api := gin.Group("/api/v1/auth")
	
	api.POST("/register", authHandler.Register)

	fmt.Println("Server started on :8080")
	gin.Run(":8080")
}
