package main

import (
	"github.com/LuigiEnzoFerrari/servers/bff/user_server/cmd/internal/handler"
	"github.com/LuigiEnzoFerrari/servers/bff/user_server/cmd/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)
	server := gin.Default()
	group := server.Group("/api/v1")
	group.GET("/users/:user_id", userHandler.GetUsersByUserID)
	server.Run(":8082")
}
