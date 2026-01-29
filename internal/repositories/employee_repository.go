package repositories

import (
	"context"

	"github.com/GowthamMuddusetty/employee-management-api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository struct {
	db *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(ctx context.Context, e *models.Employee) error {
	query := `
		INSERT INTO employees (id, full_name, job_title, country, salary)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(ctx, query, e.ID, e.FullName, e.JobTitle, e.Country, e.Salary)
	return err
}

func (r *EmployeeRepository) GetByID(ctx context.Context, id string) (*models.Employee, error) {
	query := `
		SELECT id, full_name, job_title, country, salary, created_at, updated_at
		FROM employees
		WHERE id = $1
	`

	e := &models.Employee{}
	err := r.db.QueryRow(ctx, query, id).
		Scan(&e.ID, &e.FullName, &e.JobTitle, &e.Country, &e.Salary, &e.CreatedAt, &e.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *EmployeeRepository) List(ctx context.Context) ([]*models.Employee, error) {
	query := `
		SELECT id, full_name, job_title, country, salary, created_at, updated_at
		FROM employees
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*models.Employee
	for rows.Next() {
		e := &models.Employee{}
		if err := rows.Scan(
			&e.ID, &e.FullName, &e.JobTitle,
			&e.Country, &e.Salary,
			&e.CreatedAt, &e.UpdatedAt,
		); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	return employees, nil
}

func (r *EmployeeRepository) Update(ctx context.Context, e *models.Employee) error {
	query := `
		UPDATE employees
		SET full_name = $2, job_title = $3, country = $4, salary = $5, updated_at = now()
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query, e.ID, e.FullName, e.JobTitle, e.Country, e.Salary)
	return err
}

func (r *EmployeeRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM employees WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
