package handlers

import (
	"context"
	"net/http"

	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	SignUp(ctx context.Context, password string, username string) (*domain.Auth, error)
	Login(ctx context.Context, password string, username string) (*domain.JwtToken, error)
	// Logout(ctx context.Context)
	Protected(ctx context.Context)
}

type handler struct {
	service UserUseCase
}

func NewHandler(service UserUseCase) *handler {
	return &handler{service: service}
}

type signUpRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type authLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type authLoginResponse struct {
	Token string `json:"token"`
}

func (h *handler) SignUp(c *gin.Context) {
	var req signUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.service.SignUp(c, req.Password, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user signed up"})
}

func (h *handler) Login(c *gin.Context) {
	var req authLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.service.Login(c, req.Password, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, authLoginResponse{Token: token.Token})
}

func (h *handler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user signed out"})
}

func (h *handler) Protected(c *gin.Context) {
	h.service.Protected(c)
}
