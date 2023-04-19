package postgres

import (
	"context"
	"time"

	m "github.com/osg_task/internal/controller/storage/repo"
)

var (
	create_time time.Time
	update_time time.Time
)

func (a *TaskRepo) CreateAttendance(req *m.AttendanceReq) (*m.AttendanceRes, error) {
	var res m.AttendanceRes
	err := a.db.Pool.QueryRow(context.Background(),
	`INSERT INTO 
		attendance(employee_id, type) 
	VALUES 
		($1, $2) 
	RETURNING 
		id, employee_id, type, created_at`, req.EmployeeId, req.Type).
		Scan(&res.Id, &res.EmployeeId, &res.Type, &create_time)
	if err != nil {
		return &m.AttendanceRes{}, err
	}
	res.CreatedAt = create_time.Format(time.RFC1123)
	return &res, nil
}

func (a *TaskRepo) UpdateAttendance(req *m.AttendanceUpdateReq) (*m.AttendanceRes, error) {
	var res m.AttendanceRes
	err := a.db.Pool.QueryRow(context.Background(),
	`UPDATE 
		attendance 
	SET 
		type=$1 WHERE id=$2, employee_id=$3, updated_at=NOW() 
	RETURNING 
		id, employee_id, type, created_at, updated_at`,
		req.Type, req.Id, req.EmployeeId).
		Scan(&res.Id, &res.EmployeeId, &res.Type, &create_time, &update_time)
	if err != nil {
		return &m.AttendanceRes{}, err
	}
	res.CreatedAt = create_time.Format(time.RFC1123)
	res.UpdatedAt = update_time.Format(time.RFC1123)
	return &res, nil
}
