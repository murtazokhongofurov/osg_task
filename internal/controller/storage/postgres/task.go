package postgres

import (
	"context"
	"time"

	m "github.com/osg_task/internal/controller/storage/repo"
)

var (
	createtime, updatetime    time.Time
	startedtime, finishedtime time.Time
)

func (t *TaskRepo) CreateTask(task *m.TaskReq) (*m.TaskRes, error) {
	var res m.TaskRes
	query := `
	INSERT INTO 
		tasks(developer_id, title, description, file_url, started_at, finished_at, status)
	VALUES
		($1, $2, $3, $4, $5, $6, $7)
	RETURNING 
		id, developer_id, title, description, file_url, started_at, finished_at, status, created_at`
	err := t.db.Pool.QueryRow(context.Background(), query,
		task.TeamleadId, task.Title, task.Description, task.FileUrl, task.StartedDate, task.EndDate, "todo").
		Scan(&res.Id, &res.TeamleadId, &res.Title, &res.Description, &res.FileUrl, &startedtime, &finishedtime, &res.Status, &createtime)
	if err != nil {
		return &m.TaskRes{}, err
	}
	res.CreatedAt = createtime.Format(time.RFC1123)
	res.StartedDate = startedtime.Format("2006/01/02")
	res.EndDate = finishedtime.Format("2006/01/02")
	return &res, nil
}

func (t *TaskRepo) GetTask(id int) (*m.TaskRes, error) {
	var res m.TaskRes
	query := `
	SELECT
		id, developer_id, title, description, file_url, 
		started_at, finished_at, status, created_at, updated_at
	FROM 
		tasks
	WHERE 
		id=$1`
	err := t.db.Pool.QueryRow(context.Background(), query, id).
		Scan(&res.Id, &res.TeamleadId, &res.Title, &res.Description,
			&res.FileUrl, &startedtime, &finishedtime, &res.Status, &createtime, &updatetime)
	if err != nil {
		return &m.TaskRes{}, err
	}
	res.CreatedAt = createtime.Format(time.RFC1123)
	res.UpdatedAt = updatetime.Format(time.RFC1123)
	res.StartedDate = startedtime.Format("2006/01/02")
	res.EndDate = finishedtime.Format("2006/01/02")
	return &res, nil
}

func (t *TaskRepo) UpdateTask(task *m.TaskRes) (*m.TaskRes, error) {
	var res m.TaskRes
	query := `
	UPDATE 
		tasks
	SET 
		title=$1, description=$2, file_url=$3, started_at=$4, finished_at=$5, updated_at=NOW()
	WHERE
		id=$6
	RETURNING 
		id, developer_id, title, description, file_url, started_at, finished_at, status, created_at, updated_at`
	err := t.db.Pool.QueryRow(context.Background(), query,
		task.Title, task.Description, task.FileUrl, task.StartedDate, task.EndDate, task.Id).
		Scan(&res.Id, &res.TeamleadId, &res.Title, &res.Description, &res.FileUrl,
			&startedtime, &finishedtime, &res.Status, &createtime, &updatetime)
	if err != nil {
		return &m.TaskRes{}, err
	}
	res.CreatedAt = createtime.Format(time.RFC1123)
	res.UpdatedAt = updatetime.Format(time.RFC1123)
	res.StartedDate = startedtime.Format("2006/01/02")
	res.EndDate = finishedtime.Format("2006/01/02")

	return &res, nil
}

func (t *TaskRepo) DeleteTask(id int) error {
	_, err := t.db.Pool.Exec(context.Background(), `DELETE FROM tasks WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskRepo) GetTaskDeveloperId(id string) (*m.AllTask, error) {
	var res m.AllTask
	query := `
	SELECT 
		id, title, description
	FROM 
		tasks
	WHERE 
		developer_id=$1`
	rows, err := t.db.Pool.Query(context.Background(), query, id)
	if err != nil {
		return &m.AllTask{}, err
	}
	for rows.Next() {
		temp := m.TaskRes{}
		err = rows.Scan(&temp.Id, &temp.Title, &temp.Description)
		if err != nil {
			return &m.AllTask{}, err
		}
		res.Tasks = append(res.Tasks, temp)
	}
	return &res, nil
}

func (t *TaskRepo) GetAllTask() (*m.AllTask, error) {
	var res m.AllTask
	query := `
	SELECT 
		id, developer_id, title, started_date, finished_date, status
	FROM 
		tasks`
	rows, err := t.db.Pool.Query(context.Background(), query)
	if err != nil {
		return &m.AllTask{}, err
	}
	for rows.Next() {
		temp := m.TaskRes{}
		err = rows.Scan(&temp.Id, &temp.TeamleadId, &temp.Title,
			&startedtime, &finishedtime, &temp.Status)
		if err != nil {
			return &m.AllTask{}, err
		}
		temp.StartedDate = startedtime.Format("2006/01/02")
		temp.EndDate = finishedtime.Format("2006/01/02")
		res.Tasks = append(res.Tasks, temp)
	}
	return &res, nil
}

func (t *TaskRepo) UpdateStatus(req *m.StatusUpdate) (*m.StatusUpdate, error) {
	var res m.StatusUpdate
	query := `
	UPDATE 
		tasks 
	SET 
		status=$1
	WHERE
		id=$2
	RETURNING 
		id, status`
	err := t.db.Pool.QueryRow(context.Background(), query, req.Status, req.Id).
		Scan(&res.Id, &res.Status)
	if err != nil {
		return &m.StatusUpdate{}, err
	}
	return &res, nil
}
