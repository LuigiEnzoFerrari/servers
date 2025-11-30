package main

import (
	handlers "github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/handler/http"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/middleware"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/repository"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=pass dbname=auth port=5433 sslmode=disable TimeZone=America/Sao_Paulo"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	userRepo := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepo)
	jwtService := service.NewJwtService()
	handler := handlers.NewHandler(userService)
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	engine := gin.Default()
	engine.POST("/signup", handler.SignUp)
	engine.POST("/login", handler.Login)
	engine.POST("/logout", handler.Logout)

	protected := engine.Group("/")
	protected.Use(authMiddleware.AuthMiddleware())
	protected.POST("/protected", handler.Protected)

	engine.Run(":8080")
}
