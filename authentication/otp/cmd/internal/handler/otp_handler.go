package handler

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/handler/dto"
	"github.com/gin-gonic/gin"
)

type OptService interface {
	SendOTPEmail(ctx context.Context, event domain.Event) error
	VerifyOTP(ctx context.Context, email string, otpCode string) error
}

type OptHandler struct {
	optService OptService
}

func NewOptHandler(optService OptService) *OptHandler {
	return &OptHandler{optService: optService}
}

func (h *OptHandler) VerifyOTP(c *gin.Context) {
	log, _ := c.Get("logger")
	logger := log.(*slog.Logger)
	var req dto.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.optService.VerifyOTP(c.Request.Context(), req.Username, req.Code)
	if err != nil {
		if errors.Is(err, domain.OTPNotFoundError) {
			logger.Error("OTP not found", "error", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "OTP not found"})
			return
		} else if errors.Is(err, domain.InvalidOTPError) {
			logger.Error("Invalid OTP", "error", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
			return
		}
		logger.Error("failed to verify OTP", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

