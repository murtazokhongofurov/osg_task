package repo

type Employee struct {
	Id           string
	FullName     string
	ProfilePhoto string
	BirthDate    string
	Phone        string
	Position     string
	Role         string
	RefreshToken string
}

type AllEmployee struct {
	Employees []Employee
}



type Developer struct {
	Id            string
	EmployeeId    string
	DeveloperRole string
}

type AllDeveloper struct {
	Developers []Developer
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
