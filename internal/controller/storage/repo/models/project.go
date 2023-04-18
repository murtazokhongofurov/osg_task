package models

type ProjectReq struct {
	DeveloperId  string
	Name         string
	StartedDate  string
	FinishedDate string
	Status       string
	FileUrl      string
}

type ProjectRes struct {
	Id           int
	DeveloperId  string
	Name         string
	StartedDate  string
	FinishedDate string
	Status       string
	FileUrl      string
}
