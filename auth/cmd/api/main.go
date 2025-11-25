package main

import (
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/handler/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/repository"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/service"
)

func main() {

	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=pass dbname=auth port=5433 sslmode=disable TimeZone=America/Sao_Paulo"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)


	handler := handlers.NewHandler(userService)

	engine := gin.Default()

	engine.POST("/signup", handler.SignUp)
	engine.POST("/login", handler.Login)
	engine.POST("/logout", handler.Logout)
	engine.POST("/refresh", handler.Refresh)

	engine.Run(":8080")
}
