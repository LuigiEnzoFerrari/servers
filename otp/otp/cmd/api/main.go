package main

import (
	"log"

	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/consumer"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func main() {

	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal("nats connection error: ", err)
	}
	service := service.NewOptService()

	otp := consumer.NewOtpConsumer(nc, service)
	otp.Start()
	
	r := gin.Default()
	r.Run(":8081")
}
