package repo

type TaskReq struct {
	TeamleadId string
	Title       string
	Description string
	FileUrl     string
	StartedDate string
	EndDate     string
	Status      string
}

type TaskRes struct {
	Id          int
	TeamleadId string
	Title       string
	Description string
	FileUrl     string
	StartedDate string
	EndDate     string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

type AllTask struct {
	Tasks []TaskRes
}

type StatusUpdate struct {
	Id     int
	Status string
}
