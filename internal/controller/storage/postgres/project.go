package postgres

import (
	"context"
	"time"

	m "github.com/osg_task/internal/controller/storage/repo"
)

var (
	started_date, finished_date time.Time
)

func (p *TaskRepo) CreateProject(project *m.ProjectReq) (*m.ProjectRes, error) {
	var res m.ProjectRes
	query := `
	INSERT INTO 
		projects(employee_id, project_name, started_date, finished_date, status, file_url)
	VALUES
		($1, $2, $3, $4, $5, $6)
	RETURNING 
		id, employee_id, project_name, started_date, finished_date, status, file_url
	`
	err := p.db.Pool.QueryRow(context.Background(), query,
		project.DeveloperId, project.ProjectName, project.StartedDate,
		project.FinishedDate, project.Status, project.FileUrl).
		Scan(&res.Id, &res.DeveloperId, &res.ProjectName,
			&started_date, &finished_date, &res.Status, &res.FileUrl)
	if err != nil {
		return &m.ProjectRes{}, err
	}
	res.StartedDate = startedtime.Format("2006/01/02")
	res.FinishedDate = finished_date.Format("2006/01/02")
	return &res, nil
}

func (p *TaskRepo) GetProject(id int) (*m.ProjectRes, error) {
	var res m.ProjectRes
	query := `
	SELECT 
		id, employee_id, project_name, started_date, finished_date, status, file_url
	FROM 
		projects
	WHERE 
		id=$1`
	err := p.db.Pool.QueryRow(context.Background(), query, id).
		Scan(&res.Id, &res.DeveloperId, &res.ProjectName,
			&started_date, &finished_date, &res.Status, &res.FileUrl)
	if err != nil {
		return &m.ProjectRes{}, err
	}
	res.StartedDate = started_date.Format("2006/01/02")
	res.FinishedDate = finished_date.Format("2006/01/02")
	return &res, nil
}

func (p *TaskRepo) GetAllProject() (*m.AllProject, error) {
	var res m.AllProject
	query := `
	SELECT 
		id, employee_id, project_name, started_date, finished_date, status, file_url
	FROM 
		projects`
	rows, err := p.db.Pool.Query(context.Background(), query)
	if err != nil {
		return &m.AllProject{}, err
	}
	for rows.Next() {
		var temp m.ProjectRes
		err = rows.Scan(&temp.Id, &temp.DeveloperId, &temp.ProjectName,
			&started_date, &finished_date, &temp.Status, &temp.FileUrl)
		if err != nil {
			return &m.AllProject{}, err
		}
		temp.StartedDate = started_date.Format("2006/01/02")
		temp.FinishedDate = finished_date.Format("2006/01/02")
		res.Projects = append(res.Projects, temp)
	}
	return &res, nil
}

func (p *TaskRepo) UpdateProject(project *m.ProjectRes) (*m.ProjectRes, error) {
	var res m.ProjectRes
	query := `
	UPDATE 
		projects 
	SET 
		project_name=$1, started_date, finished_date, file_url, status
	RETURNING 
		id, employee_id, project_name, started_date, finished_date, status, file_url
	`
	err := p.db.Pool.QueryRow(context.Background(), query,
		project.ProjectName, project.StartedDate,
		project.FinishedDate, project.FileUrl, project.Status).
		Scan(&res.Id, &res.DeveloperId, &res.ProjectName, &started_date,
			&finished_date, &res.Status, &res.FileUrl)
	if err != nil {
		return &m.ProjectRes{}, err
	}
	res.StartedDate = started_date.Format("2006/01/02")
	res.FinishedDate = finished_date.Format("2006/01/02")
	return &res, nil
}

func (p *TaskRepo) DeleteProject(id int) error {
	_, err := p.db.Pool.Exec(context.Background(), `DELETE FROM projects WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *TaskRepo) GetProjectEmployeeId(id string) (*m.AllProject, error) {
	res := m.AllProject{}
	query := `
	SELECT 
		id, employee_id, project_name, started_date, finished_date, status, file_url
	FROM 
		projects
	WHERE 
		employee_id=$1	
		`
	rows, err := p.db.Pool.Query(context.Background(), query, id)
	if err != nil {
		return &m.AllProject{}, err
	}
	for rows.Next() {
		temp := m.ProjectRes{}
		err = rows.Scan(&temp.Id, &temp.DeveloperId, &temp.ProjectName,
			&started_date, &finished_date, &temp.Status, &temp.FileUrl)
		if err != nil {
			return &m.AllProject{}, err
		}
		temp.StartedDate = started_date.Format("2006/01/02")
		temp.FinishedDate = finished_date.Format("2006/01/02")
		res.Projects = append(res.Projects, temp)
	}
	return &res, nil
}

func (p *TaskRepo) UpdateProjectStatus(req *m.StatusUpdate) (*m.StatusUpdate, error) {
	var res m.StatusUpdate
	query := `UPDATE projects SET status=$1 WHERE id=$2 RETURNING id, status`
	err := p.db.Pool.QueryRow(context.Background(), query, req.Status, req.Id).
		Scan(&res.Id, &res.Status)
	if err != nil {
		return &m.StatusUpdate{}, err
	}
	return &res, nil
}
