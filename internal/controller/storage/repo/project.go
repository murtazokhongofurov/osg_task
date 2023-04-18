package repo

type ProjectReq struct {
	DeveloperId  string
	ProjectName  string
	StartedDate  string
	FinishedDate string
	Status       string
	FileUrl      string
}

type ProjectRes struct {
	Id           int
	DeveloperId  string
	ProjectName  string
	StartedDate  string
	FinishedDate string
	Status       string
	FileUrl      string
}

type AllProject struct {
	Projects []ProjectRes
}
