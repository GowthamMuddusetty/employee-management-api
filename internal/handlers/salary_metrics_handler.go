package handlers

import (
	"log"
	"net/http"

	"github.com/GowthamMuddusetty/employee-management-api/internal/services"
	"github.com/gin-gonic/gin"
)

type SalaryMetricsHandler struct {
	service *services.SalaryMetricsService
}

func NewSalaryMetricsHandler(service *services.SalaryMetricsService) *SalaryMetricsHandler {
	return &SalaryMetricsHandler{service: service}
}

func (h *SalaryMetricsHandler) ByCountry(c *gin.Context) {
	log.Println("Get Salary Metrics by Country endpoint called")
	country := c.Param("country")

	result, err := h.service.ByCountry(c.Request.Context(), country)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result, "message": "salary metrics fetched successfully"})
}

func (h *SalaryMetricsHandler) AvgByJobTitle(c *gin.Context) {
	log.Println("Get Average Salary by Job Title endpoint called")
	jobTitle := c.Param("jobTitle")

	result, err := h.service.AvgByJobTitle(c.Request.Context(), jobTitle)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result, "message": "average salary fetched successfully"})
}
