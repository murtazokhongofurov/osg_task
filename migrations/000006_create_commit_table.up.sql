CREATE TABLE IF NOT EXISTS comment(
    id SERIAL NOT NULL PRIMARY KEY,
    developer_id UUID REFERENCES developers(id),
    task_id INT REFERENCES tasks(id),
    text TEXT,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);