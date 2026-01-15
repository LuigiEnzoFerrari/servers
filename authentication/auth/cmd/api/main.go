package main

import (
	"log/slog"
	"os"

	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/config"
	handlers "github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/handler/http"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/infrastructure/publish"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/infrastructure/repository"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/infrastructure/security"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/service"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	config, err := config.LoadConfig()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	db, err := gorm.Open(postgres.Open(config.DNS()), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		os.Exit(1)
	}

	conn, err := amqp.Dial(config.RabbitMQURL())
	if err != nil {
		slog.Error("failed to connect to rabbitmq", "error", err)
		os.Exit(1)
	}
	ch, err := conn.Channel()
	if err != nil {
		slog.Error("failed to open rabbitmq channel", "error", err)
		os.Exit(1)
	}

	publishService, err := publish.NewRabbitMQPublish(ch)
	if err != nil {
		slog.Error("failed to create rabbitmq publisher", "error", err)
		os.Exit(1)
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
