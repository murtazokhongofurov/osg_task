package models

type EmployeeReq struct {
	FullName     string `json:"full_name"`
	ProfilePhoto string `json:"profile_photo"`
	BirthDate    string `json:"birth_date"`
	Phone        string `json:"phone"`
	Position     string `json:"position"`
	Role         string `json:"role"`
}

type EmployeeRes struct {
	Id           string `json:"id"`
	FullName     string `json:"full_name"`
	ProfilePhoto string `json:"profile_photo"`
	BirthDate    string `json:"birth_date"`
	Phone        string `json:"phone"`
	Position     string `json:"position"`
	Role         string `json:"role"`
}


