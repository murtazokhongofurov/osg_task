package postgres

import (
	"context"

	m "github.com/osg_task/internal/controller/storage/repo"
)

func (p *TaskRepo) CreateProject(project *m.ProjectReq) (*m.ProjectRes, error) {
	var res m.ProjectRes
	query := `
	INSERT INTO 
		projects(developer_id, project_name, started_date, finished_date, status, file_url)
	VALUES
		($1, $2, $3, $4, $5, $6)
	RETURNING 
		id, developer_id, project_name, started_date, finished_date, status, file_url
	`
	err := p.db.Pool.QueryRow(context.Background(), query,
		project.DeveloperId, project.ProjectName, project.StartedDate,
		project.FinishedDate, project.Status, project.FileUrl).
		Scan(&res.Id, &res.DeveloperId, &res.ProjectName,
			&res.StartedDate, &res.FinishedDate, &res.Status, &res.FileUrl)
	if err != nil {
		return &m.ProjectRes{}, err
	}
	return &res, nil
}

func (p *TaskRepo) GetProject(id int) (*m.ProjectRes, error) {
	var res m.ProjectRes
	query := `
	SELECT 
		id, developer_id, project_name, started_date, finished_date, status, file_url
	FROM 
		projects
	WHERE 
		id=$1`
	err := p.db.Pool.QueryRow(context.Background(), query, id).
		Scan(&res.Id, &res.DeveloperId, &res.ProjectName,
			&res.StartedDate, &res.FinishedDate, &res.Status, &res.FileUrl)
	if err != nil {
		return &m.ProjectRes{}, err
	}
	return &res, nil
}

func (p *TaskRepo) GetAllProject() (*m.AllProject, error) {
	var res m.AllProject
	query := `
	SELECT 
		id, developer_id, project_name, started_date, finished_date, status, file_url
	FROM 
		projects`
	rows, err := p.db.Pool.Query(context.Background(), query)
	if err != nil {
		return &m.AllProject{}, err
	}
	for rows.Next() {
		var temp m.ProjectRes
		err = rows.Scan(&temp.Id, &temp.ProjectName,
			&temp.StartedDate, &temp.FinishedDate, &temp.Status, &temp.FileUrl)
			if err != nil {
				return &m.AllProject{}, err
			}
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
		id, developer_id, project_name, started_date, finished_date, status, file_url
	`
	err := p.db.Pool.QueryRow(context.Background(), query, 
	project.ProjectName, project.StartedDate, 
	project.FinishedDate, project.FileUrl, project.Status).
	Scan(&res.Id, &res.DeveloperId, &res.ProjectName, &res.StartedDate, 
		&res.FinishedDate, &res.Status, &res.FileUrl)
	if err != nil {
		return &m.ProjectRes{}, err
	}

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
		id, developer_id, project_name, started_date, finished_date, status, file_url
	FROM 
		projects
	WHERE 
		developer_id=$1	
		`
	rows, err := p.db.Pool.Query(context.Background(), query, id)
	if err != nil {
		return &m.AllProject{}, err
	}
	for rows.Next() {
		temp := m.ProjectRes{}
		err = rows.Scan(&temp.Id, &temp.DeveloperId, &temp.ProjectName, 
			&temp.StartedDate, &temp.FinishedDate, &temp.Status, &temp.FileUrl)
		if err != nil {
			return &m.AllProject{}, err
		}
		res.Projects = append(res.Projects, temp)
	}
	return &res, nil
}