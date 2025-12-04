package main

import (
	"log"

	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/publish"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/repository"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/service"
	"github.com/nats-io/nats.go"

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

	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()
	repository := repository.NewAuthRepository(db)
	authPublish := publish.NewAuthPublish(nc)
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
