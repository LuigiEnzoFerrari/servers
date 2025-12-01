package handler

import (
	"net/http"

	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	type RegisterRequest struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authRequestDTO := domain.AuthRequestDTO{
		Email: req.Email,
		Password: req.Password,
	}

	if err := h.authService.Register(authRequestDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}