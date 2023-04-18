package storage

import (
	"github.com/osg_task/internal/controller/storage/repo"
	"github.com/osg_task/internal/pkg/db"
	"github.com/osg_task/internal/controller/storage/postgres"
)

type StorageI interface {
	Task() repo.TaskStorageI
}


type storagePg struct {
	TaskRepo repo.TaskStorageI
}

func NewStoragePg(db *db.Postgres) StorageI {
	return &storagePg{
		TaskRepo: postgres.NewTask(db),
	}
}

func (a *storagePg) Task() repo.TaskStorageI {
	return a.TaskRepo
}