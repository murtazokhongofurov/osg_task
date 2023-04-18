package postgres

import "github.com/osg_task/internal/pkg/db"

type TaskRepo struct {
	db *db.Postgres
}

func NewTask(db *db.Postgres) *TaskRepo {
	return &TaskRepo{
		db: db,
	}
}