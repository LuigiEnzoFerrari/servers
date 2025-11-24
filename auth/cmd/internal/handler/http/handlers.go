package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
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


