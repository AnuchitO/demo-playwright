package skill

import (
	"database/sql"
)

type storage struct {
	conn *sql.DB
}

func NewStorage(conn *sql.DB) *storage {
	return newStorage(conn)
}

func newStorage(conn *sql.DB) *storage {
	return &storage{
		conn: conn,
	}
}

func offset(page, limit uint) uint {
	return (page - 1) * limit
}

type storager interface {
	patcher
	GetAllSkills(orderBy string, page, limit uint) ([]Skill, uint, error)
	GetSkillByKey(key string) (Skill, error)
	CreateSkill(key string, skill record) error
	UpdateSkill(key string, skill record) error
}

func (s storage) GetSkillByKey(key string) (Skill, error) {
	row := s.conn.QueryRow("SELECT key, name, description, logo, levels, tags FROM skill WHERE key = $1", key)
	rec := &record{}
	return rec.ScanRow(row)
}

func (s storage) GetAllSkills(orderBy string, page, limit uint) ([]Skill, uint, error) {
	query := `SELECT key, name, description, logo, levels, tags, COUNT(*) OVER () AS total_items FROM skill ORDER BY $1 LIMIT $2 OFFSET $3;`
	stmt, err := s.conn.Prepare(query)
	if err != nil {
		return []Skill{}, 0, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(orderBy, limit, offset(page, limit))
	if err != nil {
		return []Skill{}, 0, err
	}
	defer rows.Close()

	return scans(rows)
}

func scans(rows *sql.Rows) ([]Skill, uint, error) {
	skills := []Skill{}
	var totalItems uint
	for rows.Next() {
		rec := &record{}
		s, err := rec.Scan(rows, &totalItems)
		if err != nil {
			return []Skill{}, 0, err
		}
		skills = append(skills, s)
	}

	return skills, totalItems, nil
}

func (s storage) CreateSkill(key string, skill record) error {
	query := `INSERT INTO skill (key, name, description, logo, levels, tags) VALUES ($1, $2, $3, $4, $5, $6);`
	stmt, err := s.conn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(key, skill.Name, skill.Description, skill.Logo, skill.Levels, skill.Tags)
	if err != nil {
		return err
	}

	return nil
}

func (s storage) UpdateSkill(key string, skill record) error {
	query := `UPDATE skill SET name = $1, description = $2, logo = $3, levels = $4, tags = $5 WHERE key = $6;`
	stmt, err := s.conn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(skill.Name, skill.Description, skill.Logo, skill.Levels, skill.Tags, key)
	if err != nil {
		return err
	}

	return nil
}
