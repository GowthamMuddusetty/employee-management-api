package handlers

import (
	"log"
	"net/http"

	"github.com/GowthamMuddusetty/employee-management-api/internal/models"
	"github.com/GowthamMuddusetty/employee-management-api/internal/services"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	service *services.EmployeeService
}

func NewEmployeeHandler(service *services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	log.Println("Create Employee endpoint called")
	var e models.Employee
	if err := c.ShouldBindJSON(&e); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.service.Create(c.Request.Context(), &e); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "employee created successfully", "id": e.ID})
}

func (h *EmployeeHandler) GetByID(c *gin.Context) {
	log.Println("Get Employee by ID endpoint called")
	id := c.Param("id")

	e, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, e)
}

func (h *EmployeeHandler) List(c *gin.Context) {
	log.Println("List Employees endpoint called")
	employees, err := h.service.List(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) Update(c *gin.Context) {
	log.Println("Update Employee endpoint called")
	id := c.Param("id")

	var e models.Employee
	if err := c.ShouldBindJSON(&e); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	e.ID = id

	if err := h.service.Update(c.Request.Context(), &e); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "employee updated successfully"})
}

func (h *EmployeeHandler) Delete(c *gin.Context) {
	log.Println("Delete Employee endpoint called")
	id := c.Param("id")

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "employee deleted successfully"})
}
