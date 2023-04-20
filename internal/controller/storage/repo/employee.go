package repo

type Employee struct {
	Id           string `json:"id"`
	FullName     string `json:"full_name"`
	ProfilePhoto string `json:"profile_photo"`
	BirthDate    string `json:"birth_date"`
	Phone        string `json:"phone"`
	Position     string `json:"position"`
	Role         string `json:"role"`
	RefreshToken string `json:"refresh_token"`
}

type AllEmployee struct {
	Employees []Employee `json:"employees"`
}

type Developer struct {
	Id            string `json:"id"`
	EmployeeId    string `json:"employee_id"`
	DeveloperRole string `json:"developer_role"`
}

type DeveloperInfo struct {
	Id            string `json:"id"`
	EmployeeId    string `json:"employee_id"`
	DeveloperRole string `json:"developer_role"`
	FullName      string `json:"full_name"`
	Position      string `json:"position"`
	ProfilePhoto  string `json:"profile_photo"`
	Phone         string `json:"phone"`
}

type AllDeveloper struct {
	Developers []DeveloperInfo `json:"developers"`
}

type CheckfieldReq struct {
	Field string
	Value string
}

type CheckfieldRes struct {
	Exists bool
}

type PhoneNumber struct {
	Phone string
}
