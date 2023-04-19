package models

type AttendanceReq struct {
	EmployeeId string `json:"employee_id"`
	Type       string `json:"type"`
}

type AttendanceRes struct {
	Id         int    `json:"id"`
	EmployeeId string `json:"employee_id"`
	Type       string `json:"type"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type AttendanceUpdateReq struct {
	Id         int    `json:"id"`
	EmployeeId string `json:"employee_id"`
	Type       string `json:"type"`
}

type AllAttendance struct {
	Attendances []AttendanceRes `json:"attendances"`
}
