package repo

import (
	m "github.com/osg_task/internal/controller/storage/repo/models"

)


type TaskStorageI interface {
	CreateUser(*m.User) (*m.User, error)
	UpdateUser(*m.User) (*m.User, error)
	GetUser(id string) (*m.User, error)
	GetAllUser() (*m.AllUser, error)
	DeleteUser(id string) error


	CreateProject()
}
