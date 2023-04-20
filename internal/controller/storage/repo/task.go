package repo

type TaskReq struct {
	TeamleadId  string
	Title       string
	Description string
	FileUrl     string
	StartedDate string
	EndDate     string
	Status      string
}

type TaskRes struct {
	Id          int    `json:"id"`
	TeamleadId  string `json:"teamlead_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FileUrl     string `json:"file_url"`
	StartedDate string `json:"started_date"`
	EndDate     string `json:"end_date"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type AllTask struct {
	Tasks []TaskRes `json:"tasks"`
}

type StatusUpdate struct {
	Id     int
	Status string
}
