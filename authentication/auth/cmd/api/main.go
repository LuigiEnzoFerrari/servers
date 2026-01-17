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
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		slog.Error("failed to open rabbitmq channel", "error", err)
		os.Exit(1)
	}
	defer ch.Close()

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

	engine := gin.New()
	engine.Use(handlers.SlogMiddleware(logger))
	engine.Use(gin.Recovery())

	api := engine.Group("/api/v1/auth")
	api.POST("/signup", handler.SignUp)
	api.POST("/login", handler.Login)
	api.POST("/logout", handler.Logout)
	api.POST("/password/forgot", handler.ForgotPassword)

	apiProtected := engine.Group("/api/v1/auth")
	apiProtected.Use(authMiddleware.Handler())
	apiProtected.Use(handlers.SlogMiddleware(logger))
	apiProtected.POST("/protected", handler.Protected)

	logger.Info("server started", "port", config.ServerPort())
	engine.Run(":" + config.ServerPort())
}
