package services

import (
	"context"

	"github.com/GowthamMuddusetty/employee-management-api/internal/repositories"
)

type CountrySalaryMetrics struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
}

type JobTitleSalaryMetrics struct {
	Avg float64 `json:"avg"`
}

type SalaryMetricsService struct {
	repo *repositories.SalaryMetricsRepository
}

func NewSalaryMetricsService(repo *repositories.SalaryMetricsRepository) *SalaryMetricsService {
	return &SalaryMetricsService{repo: repo}
}

func (s *SalaryMetricsService) ByCountry(ctx context.Context, country string) (*CountrySalaryMetrics, error) {
	min, max, avg, err := s.repo.ByCountry(ctx, country)
	if err != nil {
		return nil, err
	}

	return &CountrySalaryMetrics{Min: min, Max: max, Avg: avg}, nil
}

func (s *SalaryMetricsService) AvgByJobTitle(ctx context.Context, jobTitle string) (*JobTitleSalaryMetrics, error) {
	avg, err := s.repo.AvgByJobTitle(ctx, jobTitle)
	if err != nil {
		return nil, err
	}

	return &JobTitleSalaryMetrics{Avg: avg}, nil
}
