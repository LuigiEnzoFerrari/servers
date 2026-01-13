package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/domain"
	"github.com/gin-gonic/gin"
)

type DashboardServiceInterface interface {
	GetDashboardSummary(ctx context.Context, userID string) (*domain.DashboardSummary, error)
}

type DashboardHandler struct {
	dashboardService DashboardServiceInterface
}

func NewDashboardHandler(
	dashboardService DashboardServiceInterface,
) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: dashboardService,
	}
}

func (h *DashboardHandler) GetDashboardSummary(c *gin.Context) {
	userID := c.Param("userId")
	start := time.Now()
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}
	ctx := c.Request.Context()
	dashboardSummaryResponse, err := h.dashboardService.GetDashboardSummary(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("Time taken: ", time.Since(start))
	c.JSON(http.StatusOK, dashboardSummaryResponse)
}
