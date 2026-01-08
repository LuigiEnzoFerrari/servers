package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/LuigiEnzoFerrari/servers/bff/user_server/cmd/internal/dto"
	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	GetUsersByUserID(ctx context.Context, userID string) (*dto.GetUsersByUserIDResponse, error)
}

type UserHandler struct {
	userUseCase UserUseCase
}

func NewUserHandler(userUseCase UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) GetUsersByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	user, err := h.userUseCase.GetUsersByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, user)
}
