package models

type TaskReq struct {
	TeamleadId  string `json:"teamlead_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FileUrl     string `json:"file_url"`
	StartedDate string `json:"started_date" example:"2006/01/02"`
	EndDate     string `json:"end_date" example:"2006/01/02"`
}

type TaskRes struct {
	Id          int    `json:"id"`
	TeamleadId  string `json:"teamlead_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FileUrl     string `json:"file_url"`
	StartedDate string `json:"started_date" example:"2006/01/02"`
	EndDate     string `json:"end_date" example:"2006/01/02"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TaskUpdateReq struct {
	Id          int    `json:"id"`
	TeamleadId  string `json:"teamlead_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FileUrl     string `json:"file_url"`
	StartedDate string `json:"started_date" example:"2006/01/02"`
	EndDate     string `json:"end_date" example:"2006/01/02"`
}

type AllTask struct {
	Tasks []TaskRes `json:"tasks"`
}

type StatusUpdate struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type TaskPermissionDel struct {
	Id         int    `json:"id"`
	TeamleadId string `json:"teamlead_id"`
}
