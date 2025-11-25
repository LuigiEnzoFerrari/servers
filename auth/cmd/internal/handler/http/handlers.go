package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Service interface {
	SignUp(c *gin.Context)
	// Login(c *gin.Context)
	// Logout(c *gin.Context)
	// Refresh(c *gin.Context)
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) SignUp(c *gin.Context) {
	h.service.SignUp(c)
}

func (h *handler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user signed in"})
}

func (h *handler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user signed out"})
}

func (h *handler) Refresh(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user refreshed"})
}
