package main

import (

	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/publish"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/repository"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/service"

	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main() {

	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=pass dbname=auth port=5433 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	repository := repository.NewAuthRepository(db)
	authPublish := publish.NewAuthPublish()
	authService := service.NewAuthService(repository, authPublish)
	jwtService := service.NewJwtService()
	authHandler := handler.NewAuthHandler(authService, jwtService)

	gin := gin.Default()
	api := gin.Group("/api/v1/auth")
	
	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)
	api.POST("/forgot", authHandler.ForgotPassword)
	gin.Run(":8080")
}
