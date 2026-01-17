package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/config"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/handler"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/infrastructure/consumer"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/infrastructure/repository"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/infrastructure/smtp"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	ctx = context.WithValue(ctx, "logger", logger)
	
	defer stop()

	cfg := config.Load()
	conn, err := amqp091.Dial(cfg.RabbitMQ.URL())
	if err != nil {
		slog.Error("failed to connect to RabbitMQ", "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		slog.Error("failed to open channel", "error", err)
		os.Exit(1)
	}
	defer ch.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       0,
	})
	redisRepository := repository.NewRedisRepository(redisClient)


	smtpService := smtp.NewMailHogService(
		cfg.Smtp.Host,
		cfg.Smtp.Port,
		cfg.Smtp.Sender,
		cfg.Smtp.Password,
	)

	service := service.NewOptService(smtpService, redisRepository)

	optHandler := handler.NewOptHandler(service)

	rabbitMQConsumer, err := consumer.NewRabbitMQConsumer(ch, service, 1)
	if err != nil {
		slog.Error("failed to create RabbitMQ consumer", "error", err)
		os.Exit(1)
	}

	r := gin.Default()
	api := r.Group("/api/v1/otp")
	api.POST("/validation", optHandler.VerifyOTP)

	srv := &http.Server{
		Addr:    ":" + cfg.Server.GetPort(),
		Handler: r,
	}

	go func() {
		rabbitMQConsumer.Start(ctx, "otp.passwordforgot")
	}()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		}
	}()

	<-ctx.Done()
	
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("failed to shutdown server", "error", err)
	}
}
