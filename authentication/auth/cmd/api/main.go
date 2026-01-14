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
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/infrastructure/publish"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	config, _ := config.LoadConfig()

	db, err := gorm.Open(postgres.Open(config.DNS()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println(config.DNS())
	if err != nil {
		panic("failed to load config")
	}

	conn, err := amqp.Dial(config.RabbitMQURL())
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	publishService, err := publish.NewRabbitMQPublish(ch)
	if err != nil {
		panic(err)
	}

	authRepo := repository.NewPostgresAuthRepository(db)
	jwtService := security.NewJwtService(
		config.JwtConfig.Key,
		config.JwtConfig.ExpireTime,
	)

	authService := service.NewAuthService(authRepo, jwtService, publishService)
	middlewareService := service.NewMiddlewareService(jwtService)

	handler := handlers.NewHandler(authService)
	authMiddleware := handlers.NewAuthMiddleware(middlewareService)

	engine := gin.Default()
	group := engine.Group("/api/v1/auth")

	group.POST("/signup", handler.SignUp)
	group.POST("/login", handler.Login)
	group.POST("/logout", handler.Logout)
	group.POST("/password/forgot", handler.ForgotPassword)

	protected := engine.Group("/api/v1/auth")
	protected.Use(authMiddleware.Handler())
	protected.POST("/protected", handler.Protected)

	engine.Run(":" + config.ServerPort())
}
