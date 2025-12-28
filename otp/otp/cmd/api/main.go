package main

import (
	"context"
	"log"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/consumer"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/smtp"
)

func main() {

	conn, err := amqp091.Dial("amqp://admin:secret@localhost:5672/dev")
	log.Println("Connected to RabbitMQ")
	if err != nil {
		panic(err)
	}

	smtpService := smtp.NewMailHogService("localhost", "1025", "test@example.com", "secret")

	service := service.NewOptService(smtpService)
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
	r.Run(":8081")
}
