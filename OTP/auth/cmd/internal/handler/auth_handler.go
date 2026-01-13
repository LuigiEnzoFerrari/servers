package handler

import (
	"net/http"

	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
	jwtService *service.JwtService
}

func NewAuthHandler(authService *service.AuthService, jwtService *service.JwtService) *AuthHandler {
	return &AuthHandler{authService: authService, jwtService: jwtService}
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
	
	c.JSON(http.StatusOK, gin.H{})
}

func (h *AuthHandler) Login(c *gin.Context) {
	type LoginRequest struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.Login(req.Email, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if token, err := h.jwtService.GenerateToken(req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
	
}

func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	type ForgotPasswordRequest struct {
		Email string `json:"email" binding:"required"`
	}

	var req ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.ForgotPassword(c.Request.Context(), req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
