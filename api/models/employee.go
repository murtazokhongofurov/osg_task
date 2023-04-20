package models

type EmployeeReq struct {
	FullName     string `json:"full_name"`
	ProfilePhoto string `json:"profile_photo"`
	BirthDate    string `json:"birth_date" example:"2006/01/02"`
	Phone        string `json:"phone"`
	Position     string `json:"position"`
	Role         string `json:"role"`
}

type EmployeeRes struct {
	Id           string `json:"id"`
	FullName     string `json:"full_name"`
	ProfilePhoto string `json:"profile_photo"`
	BirthDate    string `json:"birth_date" example:"2006/01/02"`
	Phone        string `json:"phone"`
	Position     string `json:"position"`
	Role         string `json:"role"`
}

type EmployeeList struct {
	Employees []EmployeeRes `json:"employees"`
}

type GetEmployee struct {
	Id           string `json:"id"`
	FullName     string `json:"full_name"`
	ProfilePhoto string `json:"profile_photo"`
	BirthDate    string `json:"birth_date"`
	Phone        string `json:"phone"`
	Position     string `json:"position"`
	Role         string `json:"role"`
	AccessToken  string `json:"access_token"`
}

type DeveloperReq struct {
	EmployeeId    string `json:"employee_id"`
	DeveloperRole string `json:"developer_role"`
}

type DeveloperRes struct {
	Id            string `json:"id"`
	EmployeeId    string `json:"employee_id"`
	DeveloperRole string `json:"developer_role"`
	FullName      string `json:"full_name"`
	Position      string `json:"position"`
	ProfilePhoto  string `json:"profile_photo"`
}

type AllDeveloperRole struct {
	DeveloperRoles []DeveloperRes `json:"developers"`
}
