package models

type User struct {
	Id           string
	FullName     string
	ProfilePhoto string
	BirthDate    string
	Phone        string
	Position     string
	Role         string
}

type AllUser struct {
	Users []User
}
