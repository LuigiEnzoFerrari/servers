package handlers

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	SignUp(ctx context.Context, password string, username string) (*domain.Auth, error)
	Login(ctx context.Context, password string, username string) (*domain.JwtToken, error)
	Protected(ctx context.Context)
	GenerateToken(ctx context.Context, username string) (string, error)
	ForgotPassword(ctx context.Context, username string) error
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


type authLoginResponse struct {
	Token string `json:"token"`
}

func (h *handler) Login(c *gin.Context) {
	var req authLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.service.Login(c, req.Password, req.Username)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) || errors.Is(err, domain.ErrUserNotFound) {
			slog.Warn("login failed: invalid credentials", "username", req.Username, "ip", c.ClientIP())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}
		slog.Error("login failed: internal error", 
            "error", err,
            "username", req.Username,
        )
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, authLoginResponse{Token: token.Token})
}

func (h *handler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user signed out"})
}

func (h *handler) ForgotPassword(c *gin.Context) {
	type ForgotPasswordRequest struct {
		Username string `json:"username" binding:"required"`
	}

	var req ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.ForgotPassword(c.Request.Context(), req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}


func (h *handler) Protected(c *gin.Context) {
	h.service.Protected(c)
}