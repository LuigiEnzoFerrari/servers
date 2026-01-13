package handler

import (
	"net/http"
	"time"

	"github.com/LuigiEnzoFerrari/servers/bff/mock_servers/order_server/cmd/internal/handler/dto"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func (h *OrderHandler) GetOrdersByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	orders := mockOrder()
	time.Sleep(4 * time.Second)
	c.JSON(http.StatusOK, orders)
}

func mockOrder() *dto.GetOrdersByUserIDResponse {
	return &dto.GetOrdersByUserIDResponse{
		Count: 3,
		Data: []dto.Order{
			{
				OrderID:     "ord_998",
				Status:      "PENDING",
				CreatedAt:   time.Now(),
				TotalAmount: 120.50,
				Currency:    "USD",
				Items: []dto.OrderItem{
					{
						ProductID: "prod_555",
						Quantity:  2,
						UnitPrice: 60.25,
					},
				},
			},
			{
				OrderID:     "ord_887",
				Status:      "COMPLETED",
				CreatedAt:   time.Now(),
				TotalAmount: 45.00,
				Currency:    "USD",
				Items: []dto.OrderItem{
					{
						ProductID: "prod_101",
						Quantity:  1,
						UnitPrice: 45.00,
					},
				},
			},
			{
				OrderID:     "ord_665",
				Status:      "CANCELLED",
				CreatedAt:   time.Now(),
				TotalAmount: 200.00,
				Currency:    "USD",
				Items: []dto.OrderItem{
					{
						ProductID: "prod_101",
						Quantity:  1,
						UnitPrice: 100.00,
					},
					{
						ProductID: "prod_102",
						Quantity:  1,
						UnitPrice: 100.00,
					},
				},
			},
		},
	}
}
