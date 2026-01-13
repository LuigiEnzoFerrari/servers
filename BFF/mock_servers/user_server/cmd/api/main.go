package main

import (
	"github.com/LuigiEnzoFerrari/servers/bff/mock_servers/user_server/cmd/config"
	"github.com/LuigiEnzoFerrari/servers/bff/mock_servers/user_server/cmd/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.NewConfig()
	serverPort := config.Server.GetPort()
	userHandler := handler.NewUserHandler()
	server := gin.Default()
	group := server.Group("/api/v1")
	group.GET("/users/:user_id", userHandler.GetUsersByUserID)
	server.Run(":" + serverPort)
}
