package handlers

import (
	"net/http"

	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/service"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	handler_service.SignUp(c)
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user signed in"})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user signed out"})
}

func Refresh(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user refreshed"})
}
