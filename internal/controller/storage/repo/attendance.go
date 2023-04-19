package repo

type AttendanceReq struct {
	EmployeeId string
	Type       string
}

type AttendanceRes struct {
	Id         int
	EmployeeId string
	Type       string
	CreatedAt  string
	UpdatedAt  string
}

type AttendanceUpdateReq struct {
	Id         int
	EmployeeId string
	Type       string
}

type AllAttendance struct {
	Attendances []AttendanceRes
}
