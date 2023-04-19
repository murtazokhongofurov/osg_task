package repo

type TaskStorageI interface {
	CreateEmployee(e *Employee) (*Employee, error)
	UpdateEmployee(e *Employee) (*Employee, error)
	GetEmployee(id string) (*Employee, error)
	GetAllEmployee() (*AllEmployee, error)
	DeleteEmployee(id string) error
	CheckField(*CheckfieldReq) (*CheckfieldRes, error)
	GetEmployeeByPhone(*PhoneNumber) (*Employee, error)

	CreateDeveloper(*Developer) (*Developer, error)
	GetDeveloper(id string) (*Developer, error)

	CreateProject(*ProjectReq) (*ProjectRes, error)
	GetProject(id int) (*ProjectRes, error)
	GetProjectEmployeeId(id string) (*AllProject, error)
	GetAllProject() (*AllProject, error)
	UpdateProject(*ProjectRes) (*ProjectRes, error)
	DeleteProject(id int) error
	UpdateProjectStatus(*StatusUpdate) (*StatusUpdate, error)

	CreateTask(*TaskReq) (*TaskRes, error)
	GetTask(id int) (*TaskRes, error)
	GetTaskDeveloperId(id string) (*AllTask, error)
	GetAllTask() (*AllTask, error)
	UpdateTask(*TaskRes) (*TaskRes, error)
	DeleteTask(id int) error
	UpdateStatus(*StatusUpdate) (*StatusUpdate, error)

	CreateAttendance(*AttendanceReq) (*AttendanceRes, error)
	UpdateAttendance(*AttendanceUpdateReq) (*AttendanceRes, error)

	CreateComment(*CommentReq) (*CommentRes, error)
	GetComments(id int) (*AllComment, error)
	UpdateComment(*CommentRes) (*CommentRes, error)
	DeleteComment(id int) error
}
