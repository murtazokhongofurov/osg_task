package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	m "github.com/osg_task/internal/controller/storage/repo"
)

var birthdate *time.Time

func (e *TaskRepo) CreateEmployee(employee *m.Employee) (*m.Employee, error) {
	var res m.Employee
	query := `
	INSERT INTO 
		employees(
				id, full_name, profile_photo, phone, 
				birth_date, role, refresh_token, position
				)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING 
		id, full_name, profile_photo, phone, 
		birth_date, role, position
	`
	err := e.db.Pool.QueryRow(context.Background(), query,
		employee.Id, employee.FullName, employee.ProfilePhoto, employee.Phone,
		employee.BirthDate, employee.Role, employee.RefreshToken, employee.Position).
		Scan(&res.Id, &res.FullName, &res.ProfilePhoto, &res.Phone,
			&birthdate, &res.Role, &res.Position)
	if err != nil {
		return &m.Employee{}, err
	}
	res.BirthDate = birthdate.Format("2006/01/02")

	return &res, nil
}

func (e *TaskRepo) GetEmployee(id string) (*m.Employee, error) {
	var res m.Employee
	query := `
	SELECT 
		id, full_name, profile_photo, phone, 
		birth_date, role, position
	FROM 
		employees
	WHERE 
		id=$1
	`
	err := e.db.Pool.QueryRow(context.Background(), query, id).
		Scan(&res.Id, &res.FullName, &res.ProfilePhoto,
			&res.Phone, &birthdate, &res.Role, &res.Position)
	if err != nil {
		return &m.Employee{}, err
	}
	res.BirthDate = birthdate.Format(time.RFC1123)

	return &res, nil
}

func (e *TaskRepo) UpdateEmployee(employ *m.Employee) (*m.Employee, error) {
	var res m.Employee
	query := `
	UPDATE 
		employees
	SET
		full_name=$1, profile_photo=$2, phone=$3, birth_date=$4, position=$5, role=$6
	WHERE
		id=$7
	RETURNING 
		id, full_name, profile_photo, phone, birth_date, role, position
	`
	err := e.db.Pool.QueryRow(context.Background(), query,
		employ.FullName, employ.ProfilePhoto, employ.Phone,
		employ.BirthDate, employ.Position, employ.Role, employ.Id,
	).Scan(&res.Id, &res.FullName, &res.ProfilePhoto, &res.Phone,
		&birthdate, &res.Role, &res.Position)
	if err != nil {
		return &m.Employee{}, err
	}
	res.BirthDate = birthdate.Format(time.RFC1123)

	return &res, nil
}

func (e *TaskRepo) GetAllEmployee() (*m.AllEmployee, error) {
	var res m.AllEmployee

	query := `
	SELECT 
		id, full_name, profile_photo, role, position
	FROM 
		employees`
	rows, err := e.db.Pool.Query(context.Background(), query)

	if err != nil {
		return &m.AllEmployee{}, err
	}
	for rows.Next() {
		temp := m.Employee{}
		err = rows.Scan(&temp.Id, &temp.FullName, &temp.ProfilePhoto, &temp.Role, &temp.Position)
		if err != nil {
			return &m.AllEmployee{}, err
		}
		res.Employees = append(res.Employees, temp)
	}

	return &res, nil
}

func (e *TaskRepo) DeleteEmployee(id string) error {
	_, err := e.db.Pool.Exec(context.Background(), `DELETE FROM employees WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (e *TaskRepo) CheckField(req *m.CheckfieldReq) (*m.CheckfieldRes, error) {
	res := &m.CheckfieldRes{Exists: false}
	query := fmt.Sprintf("SELECT 1 FROM admins WHERE %s=$1", req.Field)
	var temp = 0
	err := e.db.Pool.QueryRow(context.Background(), query, req.Value).Scan(&temp)
	if err == sql.ErrNoRows {
		return res, nil
	} else if err != nil {
		return res, nil
	}
	if temp == 1 {
		res.Exists = true
		return res, nil
	}
	return res, nil
}

func (e *TaskRepo) GetEmployeeByPhone(req *m.PhoneNumber) (*m.Employee, error) {
	var res m.Employee
	query := `
	SELECT 
		id, full_name, profile_photo, phone, 
		birth_date, role, position	
	FROM 
		employees
	WHERE
		phone=$1`
	err := e.db.Pool.QueryRow(context.Background(), query, req.Phone).
		Scan(&res.Id, &res.FullName, &res.ProfilePhoto, &res.Phone, &birthdate, &res.Role, &res.Position)
	if err != nil {
		return &m.Employee{}, err
	}
	res.BirthDate = birthdate.Format("2006/01/02")
	return &res, nil
}

func (e *TaskRepo) CreateDeveloper(develop *m.Developer) (*m.Developer, error) {
	var res m.Developer
	query := `
	INSERT INTO 
		developers(id, employee_id, developer_role)
	VALUES
		($1, $2, $3)
	RETURNING 
		id, employee_id, developer_role
	`
	err := e.db.Pool.QueryRow(context.Background(), query,
		develop.Id, develop.EmployeeId, develop.DeveloperRole).
		Scan(&res.Id, &res.EmployeeId, &res.DeveloperRole)
	if err != nil {
		return &m.Developer{}, err
	}
	return &res, nil
}

func (e *TaskRepo) GetDeveloper(id string) (*m.Developer, error) {
	var res m.Developer
	query := `
	SELECT 
		id, employee_id, developer_role
	FROM 
		developers
	WHERE
		id=$1`
	err := e.db.Pool.QueryRow(context.Background(), query, id).
		Scan(&res.Id, &res.EmployeeId, &res.DeveloperRole)
	if err != nil {
		return &m.Developer{}, err
	}
	return &res, nil
}
