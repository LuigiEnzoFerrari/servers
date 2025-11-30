package handlers

import (
	"net/http"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service *service.UserService
}

func NewHandler(service *service.UserService) *handler {
	return &handler{service: service}
}

func (h *handler) SignUp(c *gin.Context) {
	type requestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req requestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.service.SignUp(c, req.Password, req.Username)
}

func (h *handler) Login(c *gin.Context) {
	type requestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req requestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.service.Login(c, req.Password, req.Username)
}

func (h *handler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user signed out"})
}

func (h *handler) Protected(c *gin.Context) {
	h.service.Protected(c)
}