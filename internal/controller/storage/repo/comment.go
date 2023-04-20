package repo

type CommentReq struct {
	DeveloperId string
	TaskId      int
	Text        string
}

type CommentRes struct {
	Id          int    `json:"id"`
	DeveloperId string `json:"developer_id"`
	TaskId      int    `json:"task_id"`
	Text        string `json:"text"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type AllComment struct {
	Comments []CommentRes `json:"comments"`
}
