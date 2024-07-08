package skill

import (
	"encoding/json"

	"github.com/lib/pq"
)

type rower interface {
	Scan(dest ...any) error
}

func (s *record) ScanRow(rows rower) (Skill, error) {
	if err := rows.Scan(&s.Key, &s.Name, &s.Description, &s.Logo, &s.Levels, &s.Tags); err != nil {
		return Skill{}, err
	}

	levels, err := s.unmarshalLevel()
	if err != nil {
		return Skill{}, err
	}

	return s.toSkill(levels), nil
}

func (s *record) Scan(rows rower, totalItems *uint) (Skill, error) {
	if err := rows.Scan(&s.Key, &s.Name, &s.Description, &s.Logo, &s.Levels, &s.Tags, totalItems); err != nil {
		return Skill{}, err
	}
	levels, err := s.unmarshalLevel()
	if err != nil {
		return Skill{}, err
	}

	return s.toSkill(levels), nil
}

type record struct {
	Key         string
	Name        string
	Description string
	Logo        string
	Levels      []byte
	Tags        pq.StringArray
}

func (s record) unmarshalLevel() ([]Level, error) {
	levels := []Level{}
	err := json.Unmarshal(s.Levels, &levels)
	return levels, err
}

func (s record) toSkill(level []Level) Skill {
	return Skill{
		Key:         s.Key,
		Name:        s.Name,
		Description: s.Description,
		Logo:        s.Logo,
		Tags:        s.Tags,
		Levels:      level,
	}
}
