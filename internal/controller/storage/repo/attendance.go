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
}

type AllAttendance struct {
	Attendances []AttendanceRes
}