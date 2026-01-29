package services

import (
	"context"

	"github.com/GowthamMuddusetty/employee-management-api/internal/models"
	"github.com/GowthamMuddusetty/employee-management-api/internal/repositories"
	"github.com/google/uuid"
)

type EmployeeService struct {
	repo *repositories.EmployeeRepository
}

func NewEmployeeService(repo *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) Create(ctx context.Context, e *models.Employee) error {
	e.ID = uuid.NewString()
	return s.repo.Create(ctx, e)
}

func (s *EmployeeService) GetByID(ctx context.Context, id string) (*models.Employee, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *EmployeeService) List(ctx context.Context) ([]*models.Employee, error) {
	return s.repo.List(ctx)
}

func (s *EmployeeService) Update(ctx context.Context, e *models.Employee) error {
	return s.repo.Update(ctx, e)
}

func (s *EmployeeService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
