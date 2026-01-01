package main

import (
	"context"
	"log"

	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/consumer"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/repository"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/service"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/smtp"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/handler"
	
)

func main() {

	conn, err := amqp091.Dial("amqp://admin:secret@localhost:5672/dev")
	if err != nil {
		panic(err)
	}
	log.Println("Connected to RabbitMQ")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	smtpService := smtp.NewMailHogService("localhost", "1025", "test@example.com", "secret")

	otpRepository := repository.NewRedisOtpRepository(redisClient)

	service := service.NewOptService(smtpService, otpRepository)
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
	api := r.Group("/api/v1")
	api.POST("/otp/validation", optHandler.VerifyOTP)

	r.Run(":8081")
}
