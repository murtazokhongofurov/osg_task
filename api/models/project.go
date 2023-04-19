package models

type ProjectReq struct {
	DeveloperId  string `json:"developer_id"`
	ProjectName  string `json:"project_name"`
	StartedDate  string `json:"started_date" example:"2006/01/02"`
	FinishedDate string `json:"finished_date" example:"2006/01/02"`
	FileUrl      string `json:"file_url"`
}

type ProjectRes struct {
	Id           int    `json:"id"`
	DeveloperId  string `json:"developer_id"`
	ProjectName  string `json:"project_name"`
	StartedDate  string `json:"started_date" example:"2006/01/02"`
	FinishedDate string `json:"finished_date" example:"2006/01/02"`
	Status       string `json:"status"`
	FileUrl      string `json:"file_url"`
}

type AllProject struct {
	Projects []ProjectRes `json:"projects"`
}
