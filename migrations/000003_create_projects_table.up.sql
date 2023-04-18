CREATE TABLE IF NOT EXISTS projects(
    id SERIAL NOT NULL PRIMARY KEY,
    developer_id UUID REFERENCES developers(id),
    project_name VARCHAR(250),
    status VARCHAR(100),
    file_url TEXT,
    started_date DATE,
    finished_date DATE
);