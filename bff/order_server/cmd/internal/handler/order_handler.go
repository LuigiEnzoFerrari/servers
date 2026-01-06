package handler

import (
	"context"
	"net/http"

	"github.com/LuigiEnzoFerrari/servers/bff/order_server/cmd/internal/dto"
	"github.com/gin-gonic/gin"
)

type OrderServiceInterface interface {
	GetOrdersByUserID(ctx context.Context, userID string) (*dto.GetOrdersByUserIDResponse, error)
}

type OrderHandler struct {
	orderService OrderServiceInterface
}

func NewOrderHandler(orderService OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) GetOrdersByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	orders, err := h.orderService.GetOrdersByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
	