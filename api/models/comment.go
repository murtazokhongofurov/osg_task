package models

type CommentReq struct {
	DeveloperId string `json:"developer_id"`
	TaskId      int    `json:"task_id"`
	Text        string `json:"text"`
}

type CommentUpdateReq struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type CommentRes struct {
	Id          int    `json:"id"`
	DeveloperId string `json:"developer_id"`
	TaskId      int    `json:"task_id"`
	Text        string `json:"text"`
	CreatedAt   string `json:"created_at"`
}

type AllComment struct {
	Comments []CommentRes `json:"created_at"`
}
