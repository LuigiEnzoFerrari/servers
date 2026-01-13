package main

import (
	"github.com/LuigiEnzoFerrari/servers/bff/mock_servers/order_server/cmd/config"
	"github.com/LuigiEnzoFerrari/servers/bff/mock_servers/order_server/cmd/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	config := config.NewConfig()
	serverPort := config.Server.GetPort()
	
	orderHandler := handler.NewOrderHandler()

	server := gin.Default()
	group := server.Group("/api/v1")
	group.GET("/orders/:user_id", orderHandler.GetOrdersByUserID)

	server.Run(":" + serverPort)
}
