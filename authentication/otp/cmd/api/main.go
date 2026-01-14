package main

import (
	"context"
	"log"
	"net/http"
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()
	conn, err := amqp091.Dial(cfg.RabbitMQ.URL())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Opened channel to RabbitMQ")
	defer ch.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       0,
	})
	log.Println("Connected to Redis")
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
		panic(err)
	}

	r := gin.Default()
	api := r.Group("/api/v1/otp")
	api.POST("/validation", optHandler.VerifyOTP)

	srv := &http.Server{
		Addr:    ":" + cfg.Server.GetPort(),
		Handler: r,
	}

	go func() {
		log.Println("RabbitMQ consumer started")
		rabbitMQConsumer.Start(ctx, "otp.passwordforgot")
	}()

	go func() {
		log.Println("Server started")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("Context cancelled, waiting for server and consumer to finish...")
	
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server shutdown: %s\n", err)
	}
	log.Println("Server shutdown and RabbitMQ consumer stopped")
}
