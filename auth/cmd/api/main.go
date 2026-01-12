package main

import (
	"log"

	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/config"
	handlers "github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/handler/http"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/infrastructure/repository"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/infrastructure/security"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	config, err := config.LoadConfig()
	log.Println(config.DNS())
	if err != nil {
		panic("failed to load config")
	}


	db, err := gorm.Open(postgres.Open(config.DNS()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	authRepo := repository.NewPostgresAuthRepository(db)
	jwtService := security.NewJwtService(
		config.JwtConfig.Key,
		config.JwtConfig.ExpireTime,
	)

	authService := service.NewAuthService(authRepo, jwtService)
	middlewareService := service.NewMiddlewareService(jwtService)

	handler := handlers.NewHandler(authService)
	authMiddleware := handlers.NewAuthMiddleware(middlewareService)

	engine := gin.Default()
	engine.POST("/signup", handler.SignUp)
	engine.POST("/login", handler.Login)
	engine.POST("/logout", handler.Logout)

	protected := engine.Group("/")
	protected.Use(authMiddleware.Handler())
	protected.POST("/protected", handler.Protected)

	engine.Run(":8080")
}
