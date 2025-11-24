package main

import (
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/handler/http"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.POST("/signup", handlers.SignUp)
	engine.POST("/login", handlers.Login)
	engine.POST("/logout", handlers.Logout)
	engine.POST("/refresh", handlers.Refresh)

	engine.Run(":8080")
}
