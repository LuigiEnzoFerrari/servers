package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/dto"
	"net/http"
	
)

type DashboardServiceInterface interface {
	UpdateSomething(request *dto.UpdateSomethingRequest) (*dto.UpdateSomethingResponse, error)
	GetDashboardSummary() (*dto.DashboardSummaryResponse, error)
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

	dashboardSummaryResponse, err := h.dashboardService.GetDashboardSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dashboardSummaryResponse)
}

func (h *DashboardHandler) UpdateSomething(c *gin.Context) {
	updateSomethingRequest := dto.UpdateSomethingRequest{}
	if err := c.ShouldBindJSON(&updateSomethingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateSomethingResponse, err := h.dashboardService.UpdateSomething(&updateSomethingRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateSomethingResponse)
}
