package repo



type TaskStorageI interface {
	CreateEmployee(e *Employee) (*Employee, error)
	UpdateEmployee(e *Employee) (*Employee, error)
	GetEmployee(id string) (*Employee, error)
	GetAllEmployee() (*AllEmployee, error)
	DeleteEmployee(id string) error

	CreateDeveloper(*Developer) (*Developer, error)

	CreateProject(*ProjectReq) (*ProjectRes, error)
	GetProject(id int) (*ProjectRes, error)
	GetProjectEmployeeId(id string) (*AllProject, error)
	GetAllProject() (*AllProject, error)
	UpdateProject(*ProjectRes) (*ProjectRes, error)
	DeleteProject(id int) error

	CreateTask(*TaskReq) (*TaskRes, error) 
	GetTask(id int) (*TaskRes, error)
	GetTaskDeveloperId(id string) (*AllTask, error)
	UpdateTask(*TaskRes) (*TaskRes, error)
	DeleteTask(id int) error
}
