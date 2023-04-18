package repo

type TaskReq struct {
	DeveloperId string
	Title       string
	Description string
	FileUrl     string
	StartedDate string
	EndDate     string
	Status      string
}

type TaskRes struct {
	Id          int
	DeveloperId string
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
