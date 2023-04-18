CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL NOT NULL PRIMARY KEY,
    developer_id UUID REFERENCES developers(id),
    title VARCHAR(250),
    description TEXT,
    file_url TEXT,
    started_at DATE,
    finished_at DATE,
    status VARCHAR(100),
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);