package repo

type AttendanceReq struct {
	EmployeeId string
	Type       string
}

type AttendanceRes struct {
	Id         int    `json:"id"`
	EmployeeId string `json:"employee_id"`
	Type       string `json:"type"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type AttendanceUpdateReq struct {
	Id         int
	EmployeeId string
	Type       string
}

type AllAttendance struct {
	Attendances []AttendanceRes `json:"attendances"`
}
