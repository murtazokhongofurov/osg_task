package repo

type CommentReq struct {
	DeveloperId string
	TaskId      int
	Text        string
}

type CommentRes struct {
	Id          int
	DeveloperId string
	TaskId      int
	Text        string
	CreatedAt   string
}

type AllComment struct {
	Comments []CommentRes
}
