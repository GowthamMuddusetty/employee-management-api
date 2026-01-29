package models

import "time"

type Employee struct {
	ID        string    `db:"id" json:"id"`
	FullName  string    `db:"full_name" json:"full_name"`
	JobTitle  string    `db:"job_title" json:"job_title"`
	Country   string    `db:"country" json:"country"`
	Salary    float64   `db:"salary" json:"salary"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
