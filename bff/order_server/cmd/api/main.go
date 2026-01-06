package main

import (
	"github.com/gin-gonic/gin"
	"github.com/LuigiEnzoFerrari/servers/bff/order_server/cmd/internal/handler"
	"github.com/LuigiEnzoFerrari/servers/bff/order_server/cmd/internal/service"
)

func main() {
	orderService := service.NewOrderService()
	orderHandler := handler.NewOrderHandler(orderService)

	server := gin.Default()
	group := server.Group("/api/v1")
	group.GET("/orders/:user_id", orderHandler.GetOrdersByUserID)

	server.Run(":8081")
}
