package repo

type ProjectReq struct {
	AdminId      string
	ProjectName  string
	StartedDate  string
	FinishedDate string
	Status       string
	FileUrl      string
}

type ProjectRes struct {
	Id           int
	AdminId      string
	ProjectName  string
	StartedDate  string
	FinishedDate string
	Status       string
	FileUrl      string
}

type AllProject struct {
	Projects []ProjectRes
}
