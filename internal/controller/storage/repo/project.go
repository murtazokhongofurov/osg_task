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
	Id           int    `json:"id"`
	AdminId      string `json:"admin_id"`
	ProjectName  string `json:"project_name"`
	StartedDate  string `json:"started_date"`
	FinishedDate string `json:"finished_date"`
	Status       string `json:"status"`
	FileUrl      string `json:"file_url"`
}

type AllProject struct {
	Projects []ProjectRes `json:"projects"`
}
