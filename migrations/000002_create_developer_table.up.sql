CREATE TABLE IF NOT EXISTS developers(
   id UUID NOT NULL PRIMARY KEY, 
   employee_id UUID REFERENCES employees(id),
   developer_role VARCHAR(100)
);