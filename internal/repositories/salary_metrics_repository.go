package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SalaryMetricsRepository struct {
	db *pgxpool.Pool
}

func NewSalaryMetricsRepository(db *pgxpool.Pool) *SalaryMetricsRepository {
	return &SalaryMetricsRepository{db: db}
}

func (r *SalaryMetricsRepository) ByCountry(ctx context.Context, country string) (min, max, avg float64, err error) {
	query := `
		SELECT 
			COALESCE(MIN(salary), 0),
			COALESCE(MAX(salary), 0),
			COALESCE(AVG(salary), 0)
		FROM employees
		WHERE country = $1
	`
	err = r.db.QueryRow(ctx, query, country).Scan(&min, &max, &avg)
	return
}

func (r *SalaryMetricsRepository) AvgByJobTitle(ctx context.Context, jobTitle string) (avg float64, err error) {
	query := `
		SELECT COALESCE(AVG(salary), 0)
		FROM employees
		WHERE job_title = $1
	`
	err = r.db.QueryRow(ctx, query, jobTitle).Scan(&avg)
	return 
}
