CREATE TABLE IF NOT EXISTS employees(
    id UUID NOT NULL PRIMARY KEY,
    full_name VARCHAR(250) NOT NULL,
    profile_photo TEXT,
    phone VARCHAR(20) NOT NULL,
    birth_date DATE,
    role VARCHAR(100) NOT NULL,
    refresh_token TEXT,
    position VARCHAR(250)
);