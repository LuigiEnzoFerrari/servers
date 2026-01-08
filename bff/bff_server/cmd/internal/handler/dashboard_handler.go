package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/dto"
	"github.com/gin-gonic/gin"
)

type DashboardServiceInterface interface {
	UpdateSomething(ctx context.Context, request *dto.UpdateSomethingRequest) (*dto.UpdateSomethingResponse, error)
	GetDashboardSummary(ctx context.Context, userID string) (*dto.DashboardSummaryResponse, error)
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

func (h *DashboardHandler) UpdateSomething(c *gin.Context) {
	ctx := c.Request.Context()
	updateSomethingRequest := dto.UpdateSomethingRequest{}
	if err := c.ShouldBindJSON(&updateSomethingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateSomethingResponse, err := h.dashboardService.UpdateSomething(ctx, &updateSomethingRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updateSomethingResponse)
}
