package postgres

import (
	"context"
	"time"

	m "github.com/osg_task/internal/controller/storage/repo"
)

func (c *TaskRepo) CreateComment(req *m.CommentReq) (*m.CommentRes, error) {
	var res m.CommentRes
	err := c.db.Pool.QueryRow(context.Background(),
		`INSERT INTO 
		comment(developer_id, task_id, text)
	VALUES
		($1, $2, $3)
	RETURNING 
		id, developer_id, task_id, text, created_at`, req.DeveloperId, req.TaskId, req.Text).
		Scan(&res.Id, &res.DeveloperId, &res.TaskId, &res.Text, &create_time)
	if err != nil {
		return &m.CommentRes{}, err
	}
	res.CreatedAt = create_time.Format(time.RFC1123)
	return &res, nil
}

func (c *TaskRepo) GetComments(id int) (*m.AllComment, error) {
	var res m.AllComment
	rows, err := c.db.Pool.Query(context.Background(),
		`SELECT 
		id, developer_id, task_id, text 
	FROM comment where task_id=$1`, id)
	if err != nil {
		return &m.AllComment{}, err
	}
	for rows.Next() {
		temp := m.CommentRes{}
		err = rows.Scan(&temp.Id, &temp.DeveloperId, &temp.TaskId, &temp.Text)
		if err != nil {
			return &m.AllComment{}, err
		}
		res.Comments = append(res.Comments, temp)
	}
	return &res, nil
}

func (c *TaskRepo) UpdateComment(req *m.CommentRes) (*m.CommentRes, error) {
	var res m.CommentRes
	err := c.db.Pool.QueryRow(context.Background(),
		`UPDATE comment SET text=$1 WHERE id=$2 
	RETURNING id, developer_id, task_id, text, created_at`, req.Text, req.Id).
		Scan(&res.Id, &res.DeveloperId, &res.TaskId, &res.Text, &create_time)
	if err != nil {
		return &m.CommentRes{}, err
	}
	res.CreatedAt = create_time.Format(time.RFC1123)
	return &res, nil
}

func (c *TaskRepo) DeleteComment(id int) error {
	_, err := c.db.Pool.Exec(context.Background(), `DELETE FROM comment WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}
