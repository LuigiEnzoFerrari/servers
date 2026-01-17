package handlers

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/handler/http/dto"
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

func (h *handler) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	log, _ := c.Get("logger")
	logger := log.(*slog.Logger)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("signup failed: invalid request", "error", err, "username", req.Username)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.service.SignUp(c, req.Password, req.Username)
	if err != nil {
		if errors.Is(err, domain.ErrConflict) {
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			logger.Warn("signup failed: user already exists", "username", req.Username, "ip", c.ClientIP())
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		logger.Error("signup failed: internal error", "error", err, "username", req.Username)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user signed up"})
}


func (h *handler) Login(c *gin.Context) {
	var req dto.AuthLoginRequest
	log, _ := c.Get("logger")
	logger := log.(*slog.Logger)

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("login failed: invalid request", "error", err, "username", req.Username)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.service.Login(c, req.Password, req.Username)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) || errors.Is(err, domain.ErrUserNotFound) {
			logger.Warn("login failed: invalid credentials", "username", req.Username, "ip", c.ClientIP())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}
		logger.Error("login failed: internal error", 
            "error", err,
            "username", req.Username,
            "ip", c.ClientIP(),
        )
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, dto.AuthLoginResponse{Token: token.Token})
}

func (h *handler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user signed out"})
}

func (h *handler) ForgotPassword(c *gin.Context) {
	var req dto.ForgotPasswordRequest
	log, _ := c.Get("logger")
	logger := log.(*slog.Logger)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("forgot password failed: invalid request", "error", err, "username", req.Username)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.ForgotPassword(c.Request.Context(), req.Username)

	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			logger.Warn("forgot password failed: user not found", "username", req.Username, "ip", c.ClientIP())
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		logger.Error("forgot password failed: ",
			"error", err,
			"username", req.Username,
			"ip", c.ClientIP())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}


func (h *handler) Protected(c *gin.Context) {
	h.service.Protected(c)
}