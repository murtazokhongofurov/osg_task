package storage

import "github.com/osg_task/internal/controller/storage/repo"

type StorageI interface {
	Task() repo.TaskStorageI
}

