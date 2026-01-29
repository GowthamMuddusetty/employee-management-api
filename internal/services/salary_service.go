package services

import "github.com/GowthamMuddusetty/employee-management-api/internal/models"

type SalaryBreakdown struct {
	Gross     float64 `json:"gross"`
	Deduction float64 `json:"deduction"`
	Net       float64 `json:"net"`
}

func CalculateSalary(e *models.Employee) SalaryBreakdown {
	var deductionRate float64

	switch e.Country {
	case "India":
		deductionRate = 0.10
	case "United States":
		deductionRate = 0.12
	default:
		deductionRate = 0
	}

	deduction := e.Salary * deductionRate

	return SalaryBreakdown{
		Gross:     e.Salary,
		Deduction: deduction,
		Net:       e.Salary - deduction,
	}
}
