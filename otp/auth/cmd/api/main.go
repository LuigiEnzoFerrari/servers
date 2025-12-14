package main

import (
	"fmt"
	"log"

	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/publish"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/repository"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/service"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/internal/config"

	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{})
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

	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Starting server on %s", serverAddr)
	if err := gin.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
