package main

import (
	"context"
	"log"

	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/config"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/consumer"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/handler"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/infrastructure/repository"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/service"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/infrastructure/smtp"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

func main() {

	cfg := config.Load()
	conn, err := amqp091.Dial(cfg.RabbitMQ.URL())
	if err != nil {
		panic(err)
	}
	log.Println("Connected to RabbitMQ")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       0,
	})

	smtpService := smtp.NewMailHogService(
		cfg.Smtp.Host,
		cfg.Smtp.Port,
		cfg.Smtp.Sender,
		cfg.Smtp.Password,
	)

	redisRepository := repository.NewRedisRepository(redisClient)

	service := service.NewOptService(smtpService, redisRepository)
	optHandler := handler.NewOptHandler(service)
	rabbitMQConsumer := consumer.NewRabbitMQConsumer(conn, []consumer.ConsumerConfig{
		{
			QueueName:   "otp.passwordforgot",
			WorkerCount: 1,
			Handler:     service.SendOTPEmail,
		},
	})
	ctx := context.Background()
	rabbitMQConsumer.Start(ctx)

	r := gin.Default()
	api := r.Group("/api/v1/otp")
	api.POST("/validation", optHandler.VerifyOTP)

	r.Run(":" + cfg.Server.GetPort())
}
