package handlers

import (
	"log"
	"net/http"

	"github.com/GowthamMuddusetty/employee-management-api/internal/services"
	"github.com/gin-gonic/gin"
)

type SalaryHandler struct {
	employeeService *services.EmployeeService
}

func NewSalaryHandler(employeeService *services.EmployeeService) *SalaryHandler {
	return &SalaryHandler{employeeService: employeeService}
}

func (h *SalaryHandler) GetSalary(c *gin.Context) {
	log.Println("Get Salary endpoint called")
	id := c.Param("id")

	e, err := h.employeeService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	result := services.CalculateSalary(e)

	c.JSON(http.StatusOK, gin.H{"message": "salary calculated successfully", "data": result})
}
