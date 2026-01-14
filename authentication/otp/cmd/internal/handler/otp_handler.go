package handler

import (
	"net/http"

	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/domain"
	"github.com/gin-gonic/gin"
)

type OptHandler struct {
	optService domain.OptService
}

func NewOptHandler(optService domain.OptService) *OptHandler {
	return &OptHandler{optService: optService}
}

func (h *OptHandler) VerifyOTP(c *gin.Context) {
	type VerifyOTPRequest struct {
		Username string `json:"username" binding:"required"`
		Code string `json:"code" binding:"required"`
	}

	var req VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.optService.VerifyOTP(c.Request.Context(), req.Username, req.Code); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

