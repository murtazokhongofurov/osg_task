CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL NOT NULL PRIMARY KEY,
    employee_id UUID REFERENCES employees(id),
    title VARCHAR(250),
    description TEXT,
    file_url TEXT,
    started_date DATE,
    finished_date DATE,
    status VARCHAR(100),
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);