package data

type Project struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	UserId     int64  `json:"userId" db:"user_id"`
	DBName     string `json:"dbName" db:"db_name"`
	JsonFields string `json:"jsonFields" db:"json_fields"`
}

func NewProject(name string, userId int64, jsonField string) *Project {
	return &Project{
		0,
		name,
		userId,
		"",
		jsonField,
	}
}

func (p *Project) CreateProject() error {
	_, err := DB.Exec("INSERT INTO project (name, user_id, db_name, json_fields) VALUES ($1, $2, $3, $4)", p.Name, p.UserId, p.DBName, p.JsonFields)

	return err
}

func (p *Project) GetProjectById() error {
	err := DB.Get(p, "SELECT * FROM project WHERE id = $1", p.Id)

	return err
}

func (p *Project) GetProjectByUserId() error {
	err := DB.Get(p, "SELECT * FROM project WHERE user_id = $1", p.UserId)

	return err
}

func (p *Project) SaveProject() error {
	_, err := DB.Exec("UPDATE project SET name = $1, json_fields = $2 WHERE id = $3", p.Name, p.JsonFields, p.Id)

	return err
}
