CREATE TABLE IF NOT EXISTS employees (
    id UUID PRIMARY KEY,
    full_name TEXT NOT NULL,
    job_title TEXT NOT NULL,
    country TEXT NOT NULL,
    salary NUMERIC NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
