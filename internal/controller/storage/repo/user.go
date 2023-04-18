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
